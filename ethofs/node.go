package ethofs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	config "github.com/ipfs/go-ipfs-config"
	files "github.com/ipfs/go-ipfs-files"
	libp2p "github.com/ipfs/go-ipfs/core/node/libp2p"
	icore "github.com/ipfs/interface-go-ipfs-core"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"

	assets "github.com/ipfs/go-ipfs/assets"
	namesys "github.com/ipfs/go-ipfs/namesys"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	// This package is needed so that all the preloaded plugins are loaded automatically
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	"github.com/libp2p/go-libp2p-core/peer"

	corehttp "github.com/ipfs/go-ipfs/core/corehttp"
	sockets "github.com/libp2p/go-socket-activation"
	manet "github.com/multiformats/go-multiaddr-net"
)

const (
	nBitsForKeypairDefault = 2048
	bitsOptionName         = "bits"
	emptyRepoOptionName    = "empty-repo"
	profileOptionName      = "profile"
)

var errRepoExists = errors.New(`ipfs configuration file already exists!
Reinitializing would overwrite your keys.
`)

// Setting up the ethoFS/IPFS Repo
func setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("Error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("Error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("Error injecting plugins: %s", err)
	}

	return nil
}

func createTempRepo(ctx context.Context) (string, error) {
	repoPath, err := ioutil.TempDir("", "ipfs-shell")
	if err != nil {
		return "", fmt.Errorf("Failed to get the temp dir: %s", err)
	}

	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return "", err
	}

	// Create the repo with the config
	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("Failed to initialize ephemeral node: %s", err)
	}

	return repoPath, nil
}

func swarmPeers(api icore.CoreAPI) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conns, err := api.Swarm().Peers(ctx)
	if err != nil {
		log.Error("ethoFS - peer swarming has failed")
	}

	for _, c := range conns {
		addr := c.Address().String()
		peer := c.ID().Pretty()
		log.Info("ethoFS - peer connection found", "addr", addr, "id", peer)
	}
}

// Creates an ethoFS/IPFS node and returns its coreAPI
func createNode(ctx context.Context, repoPath string) (icore.CoreAPI, *core.IpfsNode, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, nil, err
	}

	// Construct the node

	nodeOptions := &core.BuildCfg{
		Online: true,
		// This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		Routing: libp2p.DHTOption,
		// This option sets the node to be a client DHT node (only fetching records)
		// Routing: libp2p.DHTClientOption,
		Repo: repo,
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, nil, err
	}

	// Attach the Core API to the constructed node
	api, apiErr := coreapi.NewCoreAPI(node)
	return api, node, apiErr
}

// Spawns a node on the default repo location, if the repo exists
func spawnDefault(ctx context.Context) (icore.CoreAPI, *core.IpfsNode, error) {
	defaultPath := defaultDataDir + "/ethofs"

	if err := setupPlugins(defaultPath); err != nil {
		return nil, nil, err

	}

	return createNode(ctx, defaultPath)
}

func connectToPeers(ctx context.Context, ipfs icore.CoreAPI, peers []string) error {
	var wg sync.WaitGroup
	peerInfos := make(map[peer.ID]*peerstore.PeerInfo, len(peers))
	for _, addrStr := range peers {
		addr, err := ma.NewMultiaddr(addrStr)
		if err != nil {
			return err
		}
		pii, err := peerstore.InfoFromP2pAddr(addr)
		if err != nil {
			return err
		}
		pi, ok := peerInfos[pii.ID]
		if !ok {
			pi = &peerstore.PeerInfo{ID: pii.ID}
			peerInfos[pi.ID] = pi
		}
		pi.Addrs = append(pi.Addrs, pii.Addrs...)
	}

	wg.Add(len(peerInfos))
	for _, peerInfo := range peerInfos {
		go func(peerInfo *peerstore.PeerInfo) {
			defer wg.Done()
			err := ipfs.Swarm().Connect(ctx, *peerInfo)
			if err != nil {
				log.Debug("ethoFS - peer connection has failed", "node", peerInfo.ID, "message", err)
			} else {
				log.Info("ethoFS - peer connection was successful", "node", peerInfo.ID)
			}
		}(peerInfo)
	}
	wg.Wait()

	go swarmPeers(ipfs)

	return nil
}

func getUnixfsFile(path string) (files.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	st, err := file.Stat()
	if err != nil {
		return nil, err
	}

	f, err := files.NewReaderPathFile(path, file, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func getUnixfsNode(path string) (files.Node, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func applyProfiles(conf *config.Config, profiles string) error {
	if profiles == "" {
		return nil
	}

	for _, profile := range strings.Split(profiles, ",") {
		transformer, ok := config.Profiles[profile]
		if !ok {
			return fmt.Errorf("Invalid configuration profile: %s", profile)
		}

		if err := transformer.Transform(conf); err != nil {
			return err
		}
	}
	return nil
}

func doInit(out io.Writer, repoRoot string, empty bool, nBitsForKeypair int, confProfiles string, conf *config.Config) error {

	if err := checkWritable(repoRoot); err != nil {
		return err
	}

	if fsrepo.IsInitialized(repoRoot) {
		return errRepoExists
	}

	if conf == nil {
		var err error
		conf, err = config.Init(out, nBitsForKeypair)
		if err != nil {
			return err
		}
	}

	if err := applyProfiles(conf, confProfiles); err != nil {
		return err
	}

	if err := fsrepo.Init(repoRoot, conf); err != nil {
		return err
	}

	if !empty {
		if err := addDefaultAssets(out, repoRoot); err != nil {
			return err
		}
	}

	return initializeIpnsKeyspace(repoRoot)
}

func checkWritable(dir string) error {
	_, err := os.Stat(dir)
	if err == nil {
		// dir exists, make sure we can write to it
		testfile := filepath.Join(dir, "test")
		fi, err := os.Create(testfile)
		if err != nil {
			if os.IsPermission(err) {
				return fmt.Errorf("%s is not writeable by the current user", dir)
			}
			return fmt.Errorf("Unexpected error while checking writeablility of repo root: %s", err)
		}
		fi.Close()
		return os.Remove(testfile)
	}

	if os.IsNotExist(err) {
		// dir doesn't exist, check that we can create it
		return os.Mkdir(dir, 0775)
	}

	if os.IsPermission(err) {
		return fmt.Errorf("Cannot write to %s, incorrect permissions", err)
	}

	return err
}

func addDefaultAssets(out io.Writer, repoRoot string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := fsrepo.Open(repoRoot)
	if err != nil { // NB: repo is owned by the node
		return err
	}

	nd, err := core.NewNode(ctx, &core.BuildCfg{Repo: r})
	if err != nil {
		return err
	}
	defer nd.Close()

	dkey, err := assets.SeedInitDocs(nd)
	if err != nil {
		return fmt.Errorf("init: seeding init docs failed: %s", err)
	}
	log.Warn("init: seeded init docs %s", dkey)

	if _, err = fmt.Fprintf(out, "to get started, enter:\n"); err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "\n\tipfs cat /ipfs/%s/readme\n\n", dkey)
	return err
}

func initializeIpnsKeyspace(repoRoot string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := fsrepo.Open(repoRoot)
	if err != nil { // NB: repo is owned by the node
		return err
	}

	nd, err := core.NewNode(ctx, &core.BuildCfg{Repo: r})
	if err != nil {
		return err
	}
	defer nd.Close()

	return namesys.InitializeKeyspace(ctx, nd.Namesys, nd.Pinning, nd.PrivateKey)
}

func initializeEthofsRepo() error {

	empty := true
	nBitsForKeypair := nBitsForKeypairDefault

	var conf *config.Config

	profiles := "lowpower"
	//profiles := "default-networking"

	repoPath := defaultDataDir + "/ethofs"

	return doInit(os.Stdout, repoPath, empty, nBitsForKeypair, profiles, conf)
}

func configEthofsNode(node *core.IpfsNode, nodeType string) error {
	r := node.Repo
	cfg, err := r.Config()
	if err != nil {
		return err
	}

	var storageMax string
	var routingType string

	if nodeType == "sn" {
		storageMax = "16GB"
		routingType = "dhtclient"
	} else if nodeType == "mn" {
		storageMax = "36GB"
		routingType = "dht"
	} else if nodeType == "gn" {
		storageMax = "76GB"
		routingType = "dht"
		gatewayString := string("/ip4/0.0.0.0/tcp/80")
		if err := r.SetConfigKey("Addresses.Gateway", gatewayString); err != nil {
			return err
		}
		cfg.Addresses.Gateway = config.Strings{gatewayString}
	}

	if err := r.SetConfigKey("Datastore.StorageMax", storageMax); err != nil {
		return err
	}
	cfg.Datastore.StorageMax = storageMax
	if err := r.SetConfigKey("Routing.Type", routingType); err != nil {
		return err
	}
	cfg.Routing.Type = routingType

	return nil
}

func initializeGateway(node *core.IpfsNode) error {
	r := node.Repo
	cfg, err := r.Config()
	if err != nil {
		return err
	}

	listeners, err := sockets.TakeListeners("io.ipfs.gateway")
	if err != nil {
		return fmt.Errorf("serveHTTPGateway: socket activation failed: %s", err)
	}

	listenerAddrs := make(map[string]bool, len(listeners))
	for _, listener := range listeners {
		listenerAddrs[string(listener.Multiaddr().Bytes())] = true
	}

	writable := cfg.Gateway.Writable

	gatewayAddrs := cfg.Addresses.Gateway
	for _, addr := range gatewayAddrs {
		gatewayMaddr, err := ma.NewMultiaddr(addr)
		if err != nil {
			return fmt.Errorf("serveHTTPGateway: invalid gateway address: %q (err: %s)", addr, err)
		}

		if listenerAddrs[string(gatewayMaddr.Bytes())] {
			continue
		}

		gwLis, err := manet.Listen(gatewayMaddr)
		if err != nil {
			return fmt.Errorf("serveHTTPGateway: manet.Listen(%s) failed: %s", gatewayMaddr, err)
		}
		listenerAddrs[string(gatewayMaddr.Bytes())] = true
		listeners = append(listeners, gwLis)
	}

	// we might have listened to /tcp/0 - let's see what we are listing on
	gwType := "readonly"
	if writable {
		gwType = "writable"
	}

	for _, listener := range listeners {
		log.Info("ethoFS - gateway initialized successfully", "type", gwType, "addr", listener.Multiaddr())
	}

	var opts = []corehttp.ServeOption{
		corehttp.MetricsCollectionOption("gateway"),
		corehttp.HostnameOption(),
		corehttp.GatewayOption(writable, "/ipfs", "/ipns"),
		corehttp.VersionOption(),
		corehttp.CheckVersionOption(),
	}

	if cfg.Experimental.P2pHttpProxy {
		opts = append(opts, corehttp.P2PProxyOption())
	}

	if len(cfg.Gateway.RootRedirect) > 0 {
		opts = append(opts, corehttp.RedirectOption("", cfg.Gateway.RootRedirect))
	}

	errc := make(chan error)
	var wg sync.WaitGroup
	for _, lis := range listeners {
		wg.Add(1)
		go func(lis manet.Listener) {
			defer wg.Done()
			errc <- corehttp.Serve(node, manet.NetListener(lis), opts...)
		}(lis)
	}

	go func() {
		wg.Wait()
		close(errc)
	}()

	return nil
}

func initializeEthofsNodeRepo(nodeType string) error {

	ctx := context.Background()

	log.Info("ethoFS - initializing ethoFS node on default repo path")

	spawnDefault(ctx)

	initErr := initializeEthofsRepo()
	if initErr != errRepoExists && initErr != nil {
		log.Error("ethoFS - unable to initalize ethoFS repo on default path", "error", initErr)
		return initErr
	}

	return nil
}

func initializeEthofsNodeConfig(nodeType string) error {

	ctx := context.Background()

	_, node, _ := spawnDefault(ctx)

	err := configEthofsNode(node, nodeType)
	if err != nil {
		log.Warn("ethoFS - unable to set default node configuration", "error", err)
		return err
	} else {
		log.Info("ethoFS - node default configuration setup complete")
	}

	return nil
}

func initializeEthofsNode(nodeType string) (icore.CoreAPI, *core.IpfsNode) {

	log.Info("ethoFS - deploying ethoFS node")

	ctx := context.Background()

	log.Info("ethoFS - initializing ethoFS node on default repo path")
	ipfs, node, err := spawnDefault(ctx)
	if err != nil {
		log.Warn("ethoFS - unable to intialize ethoFS node on default repo path", "error", err)
		os.Exit(0)
	}

	// Setup ethoFS node config defaults
	err = configEthofsNode(node, nodeType)
	if err != nil {
		log.Warn("ethoFS - unable to set default node configuration", "error", err)
	} else {
		log.Info("ethoFS - node default configuration setup complete")
	}

	if nodeType == "gn" {
		err = initializeGateway(node)
		if err != nil {
			log.Error("ethoFS - error initializing gateway", "error", err)
			os.Exit(0)
		}
	}

	log.Info("ethoFS - node initialization complete")

	bootstrapNodes := []string{
                "/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	}

	connectToPeers(ctx, ipfs, bootstrapNodes)

	return ipfs, node
}
