package ethofs

import (
        "fmt"
        "os"
        "path/filepath"
        "context"
        "log"
        "github.com/ipsn/go-ipfs/core"
        "github.com/ipsn/go-ipfs/core/coreapi"
        corerepo "github.com/ipsn/go-ipfs/core/corerepo"
        coreiface "github.com/ipsn/go-ipfs/core/coreapi/interface"
        "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-ipfs-config"
        assets "github.com/ipsn/go-ipfs/assets"
        namesys "github.com/ipsn/go-ipfs/namesys"
        fsrepo "github.com/ipsn/go-ipfs/repo/fsrepo"
        loader "github.com/ipsn/go-ipfs/plugin/loader"
)

var CurrentBlockHeight int64
var ethofsNode *core.IpfsNode
var ethofsCoreApi coreiface.CoreAPI
//var ethofsDagBatch coreiface.DagBatch

func EthofsNode() {
    repoRoot := ".ethofs"
        _, pluginErr := loadPlugins(repoRoot)
        if pluginErr != nil {
            log.Printf("Load Plugin Error: ", pluginErr)
        }

        if !fsrepo.IsInitialized(repoRoot) {
                //JUST START NODE, ALREADY INITIALIZED
            //r, err := fsrepo.Open(repoRoot)
            doInit(repoRoot)
        }


        r, repoErr := fsrepo.Open(repoRoot)
        if repoErr != nil { // NB: repo is owned by the node
            log.Printf("Repo Open Error: ", repoErr)
        }
        ctx := context.Background()
        ethofsNode, err := core.NewNode(ctx, &core.BuildCfg{
                                             Online: true,
                                             Repo: r,
                                             ExtraOpts: map[string]bool{
                                                     "pubsub": true,
                                             },
        })
        if err != nil {
            log.Printf("Core Node Deploy Error: ", err)
        }else{
            go corerepo.PeriodicGC(ctx, ethofsNode)
            ethofsCoreApi, _ = coreapi.NewCoreAPI(ethofsNode)
            //ethofsDagBatch = ethofsCoreApi.Dag().Batch(ctx)
        }
}
func loadPlugins(repoPath string) (*loader.PluginLoader, error) {
        pluginpath := filepath.Join(repoPath, "plugins")
        var plugins *loader.PluginLoader
        ok, err := checkPermissions(repoPath)
        if err != nil {
            log.Printf("Check Permissions Error: ", err)
        }
        if !ok {
                pluginpath = ""
        }
        plugins, err = loader.NewPluginLoader(pluginpath)
        if err != nil {
                log.Printf("error loading plugins: ", err)
        }

        if err := plugins.Initialize(); err != nil {
                log.Printf("error initializing plugins: ", err)
        }

        if err := plugins.Run(); err != nil {
                log.Printf("error running plugins: ", err)
        }
        return plugins, nil
}
func checkPermissions(path string) (bool, error) {
        _, err := os.Open(path)
        if os.IsNotExist(err) {
                // repo does not exist yet - don't load plugins, but also don't fail
                return false, nil
        }
        if os.IsPermission(err) {
                // repo is not accessible. error out.
                return false, fmt.Errorf("error opening repository at %s: permission denied", path)
        }

        return true, nil
}
func doInit(repoRoot string){
    ctx, _ := context.WithCancel(context.Background())
    conf, initErr := config.Init(os.Stdout, 2048)
    if initErr != nil { // NB: repo is owned by the node
        log.Printf("ethoFS Config Init Error: ", initErr)
    }
    if repoInitErr := fsrepo.Init(repoRoot, conf); repoInitErr != nil {
        log.Printf("ethoFS Repo Init Error", repoInitErr)
    }
    r, repoOpenErr := fsrepo.Open(repoRoot)
    if repoOpenErr != nil { // NB: repo is owned by the node
        log.Printf("Repo Open Error: ", repoOpenErr)
    }
    nd, _ := core.NewNode(ctx, &core.BuildCfg{Repo: r})
    defer nd.Close()
    _, seedErr := assets.SeedInitDocs(nd)
    if seedErr != nil {
        log.Printf("init: seeding init docs failed: %s", seedErr)
    }
    initializeIpnsKeyspace(repoRoot)
}
func initializeIpnsKeyspace(repoRoot string) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    r, repoOpenErr := fsrepo.Open(repoRoot)
    if repoOpenErr != nil { // NB: repo is owned by the node
        log.Printf("ethoFS Repo Open Error", repoOpenErr)
    }

    nd, nodeErr := core.NewNode(ctx, &core.BuildCfg{Repo: r})
    if nodeErr != nil {
        log.Printf("ethoFS New Node Error", nodeErr)
    }
    defer nd.Close()

    namesys.InitializeKeyspace(ctx, nd.Namesys, nd.Pinning, nd.PrivateKey)
}
