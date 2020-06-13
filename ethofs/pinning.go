package ethofs

import (
	"fmt"
	"github.com/ipfs/go-ipfs-api"
	"math/rand"
	"time"
)

var MasterPinArray []string
var selfPins map[string]string
var PinCounts map[string]int
var LastPinMessage map[string]string
var NodePinningConsensus map[string]NodePins
var ErroredPinsMap map[string]string
var PinCount = int(0)

// SetExistingPinList scans current ethoFS pins and then adds them to local tracking
func SetExistingPinList() {
	sh := shell.NewShell("localhost:" + apiPort)
	var updatedPinMap map[string]string
	updatedPinMap = make(map[string]string)
	pins, err := sh.Pins()
	if err != nil {
		fmt.Println("Error.... Unable To Get Existing Pin List: ", err)
	}
	fmt.Println("Syncing Existing Local Content")
	for k := range pins {
		if pins[k].Type != "indirect" {
			updatedPinMap[k] = k
		}
	}
	selfPins = updatedPinMap
}

// AddPin adds new ethoFS pins to local repo after replication factor has been checked
func AddPin(pin string, fraudCheck bool) {
	// CHECK IF PIN IS IN FAILED PIN TRACKING ARRAY - IF NOT GO AHEAD/ELSE SKIP/RETURN
	if _, ok := ErroredPinsMap[pin]; ok {
	} else {
		shTemp := shell.NewShell("localhost:" + apiPort)
		shTemp.SetTimeout(10 * time.Second)
		err := shTemp.Pin(pin)
		if err != nil {
			// ADD FAILED PIN TO TRACKING ARRAY
			ErroredPinsMap[pin] = pin
			// LAUNCH GO ROUTINE TO ADD IMMEDIATE PIN
			go AddImmediatePin(pin, fraudCheck)
		} else {
			if fraudCheck {
				fmt.Println("ethoFS Pin Added: ", pin)
				if _, ok := PinCounts[pin]; ok {
					PinCounts[pin]++
				} else {
					PinCounts[pin] = 1
				}
			} else {
				if !CheckPinSize(pin) {
					RemovePin(pin)
				}
			}
		}
	}
}

// AddImmediatePin adds the pin to the local repo without checking replication factor
// mainly used for upload pinning giving us faster results
func AddImmediatePin(pin string, fraudCheck bool) {
	shTemp := shell.NewShell("localhost:" + apiPort)
	shTemp.SetTimeout(10000 * time.Second)
	err := shTemp.Pin(pin)
	if err != nil {
	} else {
		if fraudCheck {
			fmt.Println("ethoFS Immediate Pin Request Complete: ", pin)
			if _, ok := PinCounts[pin]; ok {
				PinCounts[pin]++
			} else {
				PinCounts[pin] = 1
			}
		} else {
			if !CheckPinSize(pin) {
				RemovePin(pin)
			}
		}
	}
	// REMOVE FROM TRACKING ARRAY
	delete(ErroredPinsMap, pin)
}

// RemovePin removes the pin from the local repo after checking replication factor
func RemovePin(pin string) {
	shTemp := shell.NewShell("localhost:" + apiPort)
	shTemp.SetTimeout(10 * time.Second)
	err := shTemp.Unpin(pin)
	if err != nil {
	} else {
		fmt.Println("ethoFS Pin Removed: ", pin)
		if _, ok := PinCounts[pin]; ok {
			if PinCounts[pin] > 0 {
				PinCounts[pin]--
			} else {
				PinCounts[pin] = 0
			}
		} else {
			PinCounts[pin] = 0
		}
	}
}

// GetImmediatePins looks for any immediate pinning requests via the p2p pubsub communication protocol
func GetImmediatePins() {
	time.Sleep(60 * time.Second)
	for {
		var PinningConsensusBroadcastChannel = selfNodeIDHashOnly + "ImmediatePinningChannel_alpha11"
		sh := shell.NewShell("localhost:" + apiPort)
		sub, err := sh.PubSubSubscribe(PinningConsensusBroadcastChannel)
		if err != nil {
			fmt.Println("Error.... Pinning Consensus Broadcast Channel Subscription Failure (In-Bound Messaging): ", err)
		} else {
			r, msgerr := sub.Next()
			if msgerr != nil {
				fmt.Println("Error.... Pinning Consensus Broadcast Channel Failure (Message Not Receieved): ", msgerr)
			} else if gatewayNodeFlag || masterNodeFlag {
				message := ByteSliceToString(r.Data)
				fmt.Println("Immediate Pinning Request Received - Adding Pin: " + message)
				contractPinSize := GetPinSizeFromContract(message)
				if ScrubFraudulentPins(message, contractPinSize) {
					go AddImmediatePin(message, true)
				} else {
					go AddImmediatePin(message, false)
				}
			}
		}
		randomWait := rand.Intn(20)
		time.Sleep(time.Duration(randomWait) * time.Second)
	}
}

func ArrayContains(checkString string, checkArray []string) bool {
	for _, n := range checkArray {
		if checkString == n {
			return true
		}
	}
	return false
}
