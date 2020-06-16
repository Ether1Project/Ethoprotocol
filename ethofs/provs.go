package ethofs

import (
	"context"
	//"fmt"
	//"os/exec"
	//"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ipfs/go-ipfs/core"
	cid "github.com/ipfs/go-cid"
	//peer "github.com/libp2p/go-libp2p-core/peer"
	//routing "github.com/libp2p/go-libp2p-core/routing"
)

// FindProvs is used to seek out providers of a specified ethoFS hash
func FindProvs(node *core.IpfsNode, hash string) int {

	if !node.IsOnline {
		log.Error("Unable to find providers - ethoFS node is not online")
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	swarmPeers(ctx)

	//ctx, cancel := context.WithCancel(req.Context)
	//ctx, events := routing.RegisterForQueryEvents(ctx)
	//ctx, events := routing.RegisterForQueryEvents(ctx)

	c, _ := cid.Parse(hash)

	provChan := node.Routing.FindProvidersAsync(ctx, c, 10)
	prov, ok := <-provChan
	if !ok || prov.ID == "" {
		log.Error("Expected provider. stream closed early")
		return 0
	} else {
		log.Info("ethoFS Provider Search Successful", "hash", hash, "node", prov.ID)
	}
	return 0
	/*pchan := node.Routing.FindProvidersAsync(ctx, c, 20)

	go func() {
	 	defer cancel()
		for p := range pchan {
			np := p
			routing.PublishQueryEvent(ctx, &routing.QueryEvent{
				Type:      routing.Provider,
				Responses: []*peer.AddrInfo{&np},
			})
		}
	}()
	for e := range events {
		/*if err := res.Emit(e); err != nil {
			return err
		}*/
	/*	for provider := range e.Responses {
			log.Info("ethoFS Provider Search Successful", "hash", hash, "node", provider)
		}
		return len(e.Responses)
	}

	return 0*/
	/*ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "ipfs", "dht", "findprovs", hash)
	out, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		if verboseFlag {
			fmt.Printf("Provider Search Timeout - Hash: %s\n", hash)
		}
		return 0
	}

	// Determine how many providers were found
	splitString := strings.Split(string(out), "\n")
	if err != nil {
		fmt.Printf("Providers Seatch Failure - Hash: %s Error: %s\n", hash, err)
	}

	if verboseFlag {
		fmt.Printf("Provider Search Successful - Hash: %s Count: %d\n", hash, len(splitString))
	}*/
	// Return provider count
	//return len(splitString)
	//return 1
}
