package main

import (
	"os"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// FindProvs is used to seek out providers of a specified ethoFS hash
func FindProvs(hash string) int {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, ipfsLocation + "ipfs", "dht", "findprovs", hash)
	newEnv := append(os.Environ(), "IPFS_PATH=" + ipfsRepoPath)
        cmd.Env = newEnv
	out, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		if verboseFlag {
			fmt.Printf("Provider Search Timeout - Hash: %s\n", hash)
		}
		return 0
	}

	// Determine how many providers were found
	splitString := strings.Split(string(out), "\n")
	if err != nil {
		fmt.Printf("Providers Seatch Failure - Hash: %s Error: %s\n", hash, err)
	}

	if verboseFlag {
		fmt.Printf("Provider Search Successful - Hash: %s Count: %d\n", hash, len(splitString))
	}
	// Return provider count
	return len(splitString)
}
