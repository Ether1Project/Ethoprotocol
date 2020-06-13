package ethofs

import (
	"fmt"
	"os"
	"time"
)

// OutputPeerInfo is used to log peer data
func OutputPeerInfo() {
	// PAUSE HERE
	time.Sleep(60 * time.Second)
	for {
		file, _ := os.Create("ethofsOutput.txt")
		for _, v := range NodeHealthConsensus {
			fmt.Fprintf(file, "%v", v.NodeID)
			fmt.Fprintf(file, "    Health: %v  GNEligible: %v  MNEligible: %v  SNEligible: %v \n", v.ActiveHistory, v.GNEligible, v.MNEligible, v.SNEligible)
		}
		time.Sleep(30 * time.Second) //PAUSE HERE
	}
}

// OutputPeerInfoCSV is used to log peer data in CSV format
func OutputPeerInfoCSV() {
	// PAUSE HERE
	time.Sleep(60 * time.Second)
	for {
		file, _ := os.Create("/var/www/html/ethofsOutput.csv")
		var count = 0
		var totalDifference = 0
		for _, v := range NodeHealthConsensus {
			fmt.Fprintf(file, "%v,", v.NodeID)
			fmt.Fprintf(file, "%v,", v.LatestBlock)
			fmt.Fprintf(file, "%v,", BlockHeight-v.LatestBlock)
			fmt.Fprintf(file, "%v,", v.LastHistoryUpdateBlockHeight)
			fmt.Fprintf(file, "%v,", v.PositiveCount)
			fmt.Fprintf(file, "%v,", v.NegativeCount)
			if v.GNEligible == "Yes" {
				fmt.Fprintf(file, "Gateway Node,")
			} else if v.MNEligible == "Yes" {
				fmt.Fprintf(file, "Masternode,")
			} else if v.SNEligible == "Yes" {
				fmt.Fprintf(file, "Service Node,")
			} else {
				fmt.Fprintf(file, "Failed Requirements,")
			}
			fmt.Fprintf(file, "%v\n", v.ActiveHistory)
			count++
			totalDifference = totalDifference + (v.LatestBlock - BlockHeight)
		}
		var averageDifference = totalDifference / count
		fmt.Fprintf(file, "Average Block Height Difference: %v\n", averageDifference)
		fmt.Fprintf(file, "Current Block Height: %v\n", BlockHeight)
		// PAUSE HERE
		time.Sleep(30 * time.Second)
	}
}

// OutputHealthMessage is used to log peer health data
func OutputHealthMessage() {
	// PAUSE HERE
	time.Sleep(60 * time.Second)
	for {
		file, _ := os.Create("ethofsMessageOutput.txt")
		fmt.Fprintf(file, "%v", HealthConsensusMessage)
		// PAUSE HERE
		time.Sleep(30 * time.Second)
	}
}
