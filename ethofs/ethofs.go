package ethofs

import (
       "github.com/ethereum/go-ethereum/node"
       "github.com/ethereum/go-ethereum/ethclient"
)
func EthofsBlock(ethofsNode *node.Node) {
    go EthofsNode()

    rpcClient, _ := ethofsNode.Attach()
    nodeClient := ethclient.NewClient(rpcClient)

    go EthofsSyncStart(nodeClient)
}
