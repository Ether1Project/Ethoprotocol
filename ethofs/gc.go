package ethofs

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/corerepo"

)

func gc(node *core.IpfsNode) {
	ctx,_ := context.WithTimeout(context.Background(), 100*time.Second)

	log.Info("ethoFS - garbage collection initiated")

	go func(node *core.IpfsNode, ctx context.Context) {
		err := corerepo.GarbageCollect(node, ctx)
		if err != nil {
			log.Error("ethoFS - garbage collection error", "error", err)
		} else {
			log.Info("ethoFS - garbage collection completed")
		}
	}(node, ctx)
}
