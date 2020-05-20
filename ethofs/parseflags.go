package main

import (
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"os"
	"os/exec"
	"runtime"
	"time"
	"github.com/fatih/color"
)

var PeerHighWater = "600"
var PeerLowWater = "400"
var serviceNodeFlag bool
var masterNodeFlag bool
var gatewayNodeFlag bool
var verboseFlag bool
var swarmPort string
var apiPort string
var ethoLocation string
var ipfsLocation string
var ipfsRepoPath string

// ParseFlags is the entry point for all flags
func ParseFlags() {
	version := flag.Bool("v", false, "Prints Current Version")
	flag.BoolVar(&masterNodeFlag, "masternode", false, "Start Masternode")
	flag.BoolVar(&serviceNodeFlag, "servicenode", false, "Start Service Node")
	flag.BoolVar(&gatewayNodeFlag, "gatewaynode", false, "Start Gateway Node")
	flag.BoolVar(&verboseFlag, "verbose", false, "Run ethoFS With Verbose Output")
	flag.StringVar(&swarmPort, "swarmport", "4001", "Set IPFS Swarm Port For Discovery")
	flag.StringVar(&apiPort, "apiport", "5001", "Set IPFS API Port")
	flag.StringVar(&ethoLocation, "etholocation", "", "Set ETHO geth.ipc Location")
	flag.StringVar(&ipfsLocation, "ipfslocation", "", "Set IPFS Path Location")
	flag.StringVar(&ipfsRepoPath, "ipfsrepopath", "", "Set IPFS Repo Path Location")
	flag.Parse()
	if verboseFlag {
		fmt.Println("ethoFS Port Configuration - Swarm: " + swarmPort + " API: " + apiPort)
	}
	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	} else if masterNodeFlag {
		fmt.Println("Ether-1 Masternode Started")
		v, _ := mem.VirtualMemory()
		path := "/"
		if runtime.GOOS == "windows" {
			path = "C:"
		}
		d, _ := disk.Usage(path)
		if (float64(v.Total)/float64(GB)) > float64(1.5) && (float64(d.Total)/float64(GB)) > float64(38.00) {
			fmt.Println("Ether-1 Masternode Resource Requirements Met")
		} else {
			fmt.Println("Ether-1 Masternode Resource Requirements Not Met - Not Eligible For Node Reward")
		}

		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)

		if (float64(v.Total) / float64(GB)) < float64(1.5) {
			boldRed.Println("Masternode does not have enough RAM")
		}

		if (float64(d.Total) / float64(GB)) < float64(38.00) {
			boldRed.Println("Masternode does not have enough Disk space")
		}

		InitialContractValues() //INITIAL LOOK TO GET SYNC STATUS - TO MAKE SURE ETHO CHAIN IS SYNCED

		SetPeerLimits()
		GetBootnodeContractValues()
		ConfigBootnodePeers()
		time.Sleep(5 * time.Second) //PAUSE HERE
	} else if serviceNodeFlag {
		fmt.Println("ethoFS Service Node Started")
		v, _ := mem.VirtualMemory()
		path := "/"
		if runtime.GOOS == "windows" {
			path = "C:"
		}
		d, _ := disk.Usage(path)
		if (float64(v.Total)/float64(GB)) > float64(0.75) && (float64(d.Total)/float64(GB)) > float64(18.00) {
			fmt.Println("Ether-1 Service Node Resource Requirements Met")
		} else {
			fmt.Println("Ether-1 Service Node Resource Requirements Not Met - Not Eligible For Node Reward")
		}

		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)

		if (float64(v.Total) / float64(GB)) < float64(0.75) {
			boldRed.Println("Service Node does not have enough RAM")
		}

		if (float64(d.Total) / float64(GB)) < float64(18.00) {
			boldRed.Println("Service Node does not have enough Disk Space")
		}

		InitialContractValues() //INITIAL LOOK TO GET SYNC STATUS - TO MAKE SURE ETHO CHAIN IS SYNCED

		SetPeerLimits()
		GetBootnodeContractValues()
		ConfigBootnodePeers()
		time.Sleep(5 * time.Second) //PAUSE HERE
	} else if gatewayNodeFlag {
		fmt.Println("ethoFS Gateway Node Started")
		v, _ := mem.VirtualMemory()
		path := "/"
		if runtime.GOOS == "windows" {
			path = "C:"
		}
		d, _ := disk.Usage(path)
		if (float64(v.Total)/float64(GB)) > float64(3.0) && (float64(d.Total)/float64(GB)) > float64(70.00) {
			fmt.Println("ethoFS Gateway Node Resource Requirements Met")
		} else {
			fmt.Println("ethoFS Gateway Node Resource Requirements Not Met - Not Eligible For Node Reward")
		}

		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)

		if (float64(v.Total) / float64(GB)) < float64(3.0) {
			boldRed.Println("Gateway Node does not have enough RAM")
		}

		if (float64(d.Total) / float64(GB)) < float64(70.00) {
			boldRed.Println("Gateway Node does not have enough Disk Space")
		}

		InitialContractValues() //INITIAL LOOK TO GET SYNC STATUS - TO MAKE SURE ETHO CHAIN IS SYNCED

		SetPeerLimits()
		GetBootnodeContractValues()
		ConfigBootnodePeers()
		time.Sleep(5 * time.Second) //PAUSE HERE
		go CheckGateway()
	} else {
		fmt.Println("No Node Type Specified - Exiting ethoFS")
		os.Exit(0)
	}
}

// ConfigBootnodePeers resets ipfs bootnode config so dynamic bootnodes can be imported from ETHO
func ConfigBootnodePeers() {
	if len(BootnodePeers) > 0 {
		cmd := exec.Command(ipfsLocation + "ipfs", "bootstrap", "rm", "--all")
		newEnv := append(os.Environ(), "IPFS_PATH=" + ipfsRepoPath)
        	cmd.Env = newEnv
		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("IPFS Bootnode Peers Removal Config Failure %s\n", err)
		}
		for k := range BootnodePeers {
			cmd = exec.Command(ipfsLocation + "ipfs", "bootstrap", "add", BootnodePeers[k])
			newEnv := append(os.Environ(), "IPFS_PATH=" + ipfsRepoPath)
        		cmd.Env = newEnv
			_, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("IPFS Bootnode Peers Config Failure %s\n", err)
			}
		}
	}
}

// SetPeerLimits initiates ipfs peer limit config
func SetPeerLimits() {
	cmd := exec.Command(ipfsLocation + "ipfs", "config", "--json", "Swarm.ConnMgr.LowWater", PeerLowWater)
	newEnv := append(os.Environ(), "IPFS_PATH=" + ipfsRepoPath)
        cmd.Env = newEnv
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("IPFS Peer Connection (Low) Config Failure %s\n", err)
	}
	cmd = exec.Command(ipfsLocation + "ipfs", "config", "--json", "Swarm.ConnMgr.HighWater", PeerHighWater)
	newEnv = append(os.Environ(), "IPFS_PATH=" + ipfsRepoPath)
        cmd.Env = newEnv
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("IPFS Peer Connection (High) Config Failure %s\n", err)
	}
}
