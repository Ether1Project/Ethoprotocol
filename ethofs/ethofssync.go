package ethofs

import (
       "time"
       "math/rand"
       "github.com/ethereum/go-ethereum/ethclient"
)
func EthofsSyncStart(nodeClient *ethclient.Client) {
    time.Sleep(10 * time.Second)//WAIT FOR CHAIN READY
    go PubSubListen("ethofsTestChannel")
    time.Sleep(10 * time.Second)//WAIT FOR CHAIN READY
    for {
        var sendMessage = []byte("ethoFS Connection Successful - Syncing Peers")
        PubSubSend(sendMessage, "ethofsTestChannel")
        randomWait := rand.Intn(30)
        time.Sleep(time.Duration(randomWait) * time.Second)
    }
}
