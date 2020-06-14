package ethofs

import (
	"fmt"
	"math/rand"
	"time"
)

var FraudulentPins map[string]int64

func GetPinSizeFromContract(pin string) int64 {
	// Future update will get saved pin size from contract here

	// Return arbitrary high number to allow all pins to pass
	return 9223372036854775807
}

// ScrubFraudulentPins compares pin to historically fraudulent uploads & re verifies
// the validity of these uploads
func ScrubFraudulentPins(pin string, size int64) bool {
	if val, ok := FraudulentPins[pin]; ok {
		if size < val {
			// Returns false if the pin size is recorded as too small or fraudulent
			return false
		}
		delete(FraudulentPins, pin)
		// Return true if the pin is not fraudlent
		return true
	}
	// Unable to make determination so return false
	return false
}

// CheckPinSize compares the saved pin size in the smart contract with actual pin size
func CheckPinSize(pin string) bool {
	// Gets saved upload size in order to verify the validity of the upload
	//contractPinSize := GetPinSizeFromContract(pin)

	// Gets the actual pin/data size to compare to the saved upload size
	/*objectSize, err := GetObjectSize(pin)
	if err != nil {
		return false
	}
	if float64(contractPinSize) < (float64(objectSize) * float64(0.75)) {
		// Adds the pin to fraudulent pin list
		FraudulentPins[pin] = objectSize
		return false
	}*/
	return true
}

// AnalyzePin compares the actual provider counts to the required replication factor to
// determine if local node should add the pin or remove the pin
func AnalyzePin(pin string) {
	pinCount := FindProvs(pin)
	if pinCount < repFactor {
		//if verboseFlag {
		//	fmt.Printf("Low Pin Count Detected - Pin Replication Initiated - Hash: %s Pin Count: %d\n", pin, pinCount)
		//}
		contractPinSize := GetPinSizeFromContract(pin)
		if ScrubFraudulentPins(pin, contractPinSize) {
			AddPin(pin, true)
		} else {
			AddPin(pin, false)
		}
	} else if pinCount > (3 * repFactor) {
		//if verboseFlag {
		//	fmt.Printf("High Pin Count Detected - Pin Removal Initiated - Hash: %s Pin Count: %d\n", pin, pinCount)
		//}
		RemovePin(pin)
	}
}

// ScanPinList is the entry point for ethoFS pin management
func ScanPinList() {
	time.Sleep(60 * time.Second)
	for {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		localMasterPinArray := MasterPinArray

		fmt.Printf("Analyzing ethoFS Content Pinning - Pins To Analyze: %d\n", len(localMasterPinArray))
		if len(localMasterPinArray) > 0 {
			for _, i := range r.Perm(len(localMasterPinArray)) {
				AnalyzePin(localMasterPinArray[i])
			}
			localSelfPins := selfPins

			// IF SELF PIN LIST HAS PIN WHICH IS NOT ON THE LIST, WE REMOVE THE PIN
			for j := range localSelfPins {
				if !(ArrayContains(localSelfPins[j], localMasterPinArray)) {
					RemovePin(localSelfPins[j])
				}
			}
		}
		randomWait := rand.Intn(60)
		time.Sleep(time.Duration(randomWait) * time.Second)
	}
}
