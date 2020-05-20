package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// GetIpfsStats determines the repos stats/usage
func GetIpfsStats() {
	for {
		cmd := exec.Command(ipfsLocation+"ipfs", "repo", "stat", "--human")
		newEnv := append(os.Environ(), "IPFS_PATH="+ipfsRepoPath)
		cmd.Env = newEnv
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("IPFS Stats Failure %s\n", err)
		} else {
			splitString := strings.Split(string(out), "\n")
			if len(splitString) > 1 && err == nil {
				selfUsedStorage, _ = strconv.Atoi(strings.TrimSpace(strings.Split(splitString[1], ":")[1]))
				selfMaxStorage, _ = strconv.Atoi(strings.TrimSpace(strings.Split(splitString[2], ":")[1]))
			}
		}
		if verboseFlag {
			fmt.Printf("ethoFS Used Storage: %d\n", UsedStorage)
			fmt.Printf("ethoFS Max Storage: %d\n", MaxStorage)
		}
		time.Sleep(500 * time.Second)
	}
}
