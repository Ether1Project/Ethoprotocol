package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// MiscManagement is the main garbage collection entry point
func MiscManagement() {
	for {
		IpfsGC()
		randomWait := rand.Intn(30)
		time.Sleep(time.Duration(randomWait) * time.Minute)
	}
}

// IpfsGC initiates the ipfs garbage collection to remove stale data
func IpfsGC() {
	cmd := exec.Command(ipfsLocation+"ipfs", "repo", "gc")
	newEnv := append(os.Environ(), "IPFS_PATH="+ipfsRepoPath)
	cmd.Env = newEnv
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("IPFS Garbage Collection Failure %s\n", err)
	} else {
		fmt.Println("IPFS Garbage Collection Initiated")
	}
}
