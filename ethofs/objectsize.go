package main

import (
	"os"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// GetObjectSize determines ipfs object size
func GetObjectSize(hash string) (int64, error) {
	cmd := exec.Command(ipfsLocation + "ipfs", "object", "stat", hash)
	newEnv := append(os.Environ(), "IPFS_PATH=" + ipfsRepoPath)
        cmd.Env = newEnv
	out, err := cmd.CombinedOutput()
	fmt.Println("ethoFS Hash: " + hash)
	if err != nil {
		fmt.Printf("ethoFS Object Size Failure %s\n", err)
	} else {
		splitString := strings.Split(string(out), "\n")
		if len(splitString) > 3 && err == nil {
			cummulativeSize, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(splitString[4], ":")[1]), 10, 64)
			if verboseFlag {
				fmt.Printf("ethoFS Object Cummulative Size: %d\n", cummulativeSize)
			}
			return cummulativeSize, nil
		}
	}
	return 0, errors.New("Unable To Determine IPFS Object Size")
}
