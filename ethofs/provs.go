package ethofs

import (
	"context"
	"fmt"
	//"os/exec"
	//"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"

	cid "github.com/ipfs/go-cid"
	"github.com/ipfs/go-ipfs/core"
	peer "github.com/libp2p/go-libp2p-core/peer"
	routing "github.com/libp2p/go-libp2p-core/routing"
)

// FindProvs is used to seek out providers of a specified ethoFS hash
func FindProvs(node *core.IpfsNode, hash string) (uint64, error) {

	if !node.IsOnline {
		return 0, fmt.Errorf("Unable to find providers - ethoFS node is not online")
	}

	log.Debug("ethoFS provider search has been initiated", "hash", hash)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	ctx, events := routing.RegisterForQueryEvents(ctx)

	c, _ := cid.Parse(hash)

	pchan := node.Routing.FindProvidersAsync(ctx, c, int(repFactor*2))

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
		for _, provData := range e.Responses {
			log.Debug("ethoFS - Data provider found", "hash", hash, "node", provData.ID)
		}
		log.Debug("ethoFS - Provider search has been completed", "providers", len(e.Responses), "hash", hash)
		return uint64(len(e.Responses)), nil
	}

	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("Error - ethoFS provider search was not successful")
	}
}
