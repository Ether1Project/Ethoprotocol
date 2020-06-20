package ethofs

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/log"

	cid "github.com/ipfs/go-cid"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	options "github.com/ipfs/interface-go-ipfs-core/options"
	path "github.com/ipfs/interface-go-ipfs-core/path"
)

func pinSearch(api coreiface.CoreAPI, hash string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cid, err := cid.Parse(hash)
	if err != nil {
		return false, err
	}

	opt, err := options.Pin.Ls.Type("all")
	if err != nil {
		return false, err
	}

	pins, err := api.Pin().Ls(ctx, opt)
	if err != nil {
		return false, err
	}

	for p := range pins {
		if p != nil && p.Path() != nil {
			log.Debug("ethoFS - pin found", "type", p.Type(), "hash", p.Path().Cid())
			if cid == p.Path().Cid() {
				log.Info("ethoFS - pin match found", "type", p.Type(), "hash", p.Path().Cid())
				return true, nil
			}
		}
	}
	log.Info("ethoFS - pin match not found", "hash", cid)
	return false, nil
}

func pinAdd(api coreiface.CoreAPI, hash string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cid, err := cid.Parse(hash)
	if err != nil {
		return hash, err
	}

	resolvedPath := path.IpfsPath(cid)

	if err := api.Pin().Add(ctx, resolvedPath, options.Pin.Recursive(true)); err != nil {
		return hash, err
	}

	return hash, nil
}

func pinRemove(api coreiface.CoreAPI, hash string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cid, err := cid.Parse(hash)
	if err != nil {
		return hash, err
	}

	resolvedPath := path.IpfsPath(cid)

	if err := api.Pin().Rm(ctx, resolvedPath, options.Pin.RmRecursive(true)); err != nil {
		return hash, err
	}

	return hash, nil
}
