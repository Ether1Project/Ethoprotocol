package ethofs

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/corerepo"

)

func gc(node *core.IpfsNode) {
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)

	gcChan := corerepo.GarbageCollectAsync(node, ctx)

	log.Info("ethoFS - garbage collection initiated")

	for {
		select {
			case gcResp := <-gcChan:
				log.Info("ethoFS - garbage collection response received", "response", gcResp, "channel", gcChan)
			case <-ctx.Done():
				log.Info("ethoFS - grabage collection completed")
				return
		}
	}
}
