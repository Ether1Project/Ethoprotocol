package ethofs

import (
       "context"
       "github.com/ethereum/go-ethereum/log"
)
func PubSubSend(message []byte, channel string) {
   ctx := context.Background()
   err := ethofsCoreApi.PubSub().Publish(ctx, channel, message)
   if err != nil {
       log.Warn("Failed To Send Message - PubSub")
   }
}
func PubSubListen(channel string) string {
   ctx := context.Background()
   sub, err := ethofsCoreApi.PubSub().Subscribe(ctx, channel)
   if err != nil {
       log.Warn("Failed To Subscribe To Channel - PubSub")
   }
   for{
       message, err := sub.Next(ctx)
       if err != nil {
           log.Warn("Failed To Receive Message - PubSub")
       }else{
           lastMessage := string(message.Data())
           byteArray := []byte(lastMessage)
           log.Info("Incoming ethoFS Message: " + string(byteArray))
       }
   }
}
