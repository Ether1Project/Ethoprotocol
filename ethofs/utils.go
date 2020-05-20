package main

import (
	"bytes"
	"fmt"
	"github.com/glendc/go-external-ip"
	"github.com/ipfs/go-ipfs-api"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"unsafe"
)

func ByteSliceToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

// CheckGateway determines whether or not a valid gateway is active
func CheckGateway() {
	time.Sleep(200 * time.Second)
	for {
		sh := shell.NewShell("localhost:" + apiPort)
		mhash, err := sh.Add(bytes.NewBufferString("true"))
		if err != nil {
			log.Println("Error.... ethoFS Gateway Upload Verify Fail: ", err)
		} else {
			time.Sleep(10 * time.Second)

			consensus := externalip.DefaultConsensus(nil, nil)
			ip, err := consensus.ExternalIP()
			ipAddressString := ip.String()
			if err == nil {
				fmt.Println("External IP Address: " + ipAddressString) // print IPv4/IPv6 in string format
				url := "http://" + ipAddressString + "/ipfs/" + mhash
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					fmt.Println("ethoFS Gateway Node Verification Failure - Unable To Serve Content")
				} else {
					res, err := http.DefaultClient.Do(req)
					if err != nil {
						fmt.Println("ethoFS Gateway Node Verification Failure - Unable To Serve Content")
					} else {
						defer res.Body.Close()
						body, err := ioutil.ReadAll(res.Body)
						if err != nil {
							fmt.Println("ethoFS Gateway Node Verification Failure - Unable To Serve Content")
						} else {
							ipfsResponse := string(body)
							if strings.Contains(ipfsResponse, "true") {
								gatewayTest = "Yes"
							} else {
								gatewayTest = "No"
							}
						}
					}
				}
			}
		}
		time.Sleep(5 * time.Minute)
	}
}
