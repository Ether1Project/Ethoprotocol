package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// GetIpfsId determines ipfs id
func GetIpfsId() (string, error) {
	cmd := exec.Command(ipfsLocation+"ipfs", "id", "-f=<id>")
	newEnv := append(os.Environ(), "IPFS_PATH="+ipfsRepoPath)
	cmd.Env = newEnv
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("ethoFS ID Failure %s\n", err)
	} else {
		return string(out), nil
	}
	return "", errors.New("Unable To Determine IPFS ID")
}
