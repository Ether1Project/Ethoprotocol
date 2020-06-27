package ethofs

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"math/rand"
	"io"
	"time"

	"github.com/ipfs/go-ipfs-files"
	cid "github.com/ipfs/go-cid"
	path "github.com/ipfs/interface-go-ipfs-core/path"
)

func updatePinContractValues() error {
	c, err := ethclient.Dial(ipcLocation)
	if err != nil {
		return err
	}

	address := common.HexToAddress("0xD3b80c611999D46895109d75322494F7A49D742F")
	contract, err := NewPinStorage(address, c)
	if err != nil {
		return err
	}
	repFactorResp, err := contract.ReplicationFactor(nil)
	if err != nil {
		return err
	}
	repFactor = uint64(repFactorResp)

	pinCountResp, err := contract.PinCount(nil)
	if err != nil {
		return err
	}

	lowerRange := rand.Intn(int(pinCountResp))
	for j := uint64(0); j < uint64(30); j++ {

		go func(x uint64) {

			var pinNumber uint64

			if uint64(lowerRange) + x >= uint64(pinCountResp) {
				pinNumber = (uint64(lowerRange) + x) - uint64(pinCountResp)
			} else {
				pinNumber = uint64(lowerRange) + x
			}
			i := new(big.Int).SetUint64(pinNumber)
			contractPin, err := contract.Pins(nil, i)
			if err != nil {
				log.Debug("ethoFS - ether-1 contract connection error (contract pin)", "error",  err, "number", i)
				return
			}

			cid, err := cid.Parse(contractPin)
        		if err != nil {
                		log.Debug("ethoFS - ether-1 contract connection error (cid parse)", "error",  err)
				return
			}
			// Request serialized pin list stored on ethoFS
			resolvedPath := path.IpfsPath(cid)

			ctx, cancelCtx := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancelCtx()

			resp, err := Ipfs.Unixfs().Get(ctx, resolvedPath)
			if err != nil {
				log.Debug("ethoFS - data retrieval error", "hash", cid, "error", err)
				return
			}
			var file files.File
			file = files.ToFile(resp)
			var r io.Reader = file
			if r == nil {
				return
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(r)
			cids := scanForCids(buf.Bytes())

			for _,pin := range cids {
				log.Debug("ethoFS - pin request detail", "hash", pin, "number", i)
				pinned := pinSearch(pin)
				if !pinned {
					log.Debug("ethoFS - pin search error", "error", "pin not found")
					continue
				} else {
					log.Debug("ethoFS - data is pinned to local node", "hash", pin)
				}

				providerCount, err := FindProvs(Node, pin)
				if err != nil {
					log.Warn("ethoFS - provider search error", "error", err)
					continue
				}

				// For testing
				//log.Info("ethoFS - replication factor comparisons", "low", (repFactor / uint64(2)), "high", (repFactor + (repFactor / uint64(2))))

				if !pinned && providerCount < (repFactor / uint64(2))  {
					// Pin data due to insufficient existing providers
					addedPin, err := pinAdd(Ipfs, pin)
					if err != nil {
						log.Debug("ethoFS - pin add error", "hash", pin, "error", err)
						continue
					} else {
						log.Debug("ethoFS - pin add successful", "hash", addedPin)
					}
				} else if pinned && providerCount > (repFactor + (repFactor / uint64(2)))  {
					// Pin data due to insufficient existing providers
					removedPin, err := pinRemove(Ipfs, pin)
					if err != nil {
						log.Debug("ethoFS - pin remove error", "hash", pin, "error", err)
						continue
					} else {
						log.Debug("ethoFS - pin removal successful", "hash", removedPin)
					}
				}
			}
		}(j)
	}

	return nil
}
