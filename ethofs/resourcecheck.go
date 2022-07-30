package ethofs

import (
	"os"
	"runtime"

	"github.com/ethereum/go-ethereum/log"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func checkResources(nodeType string) {

	v, _ := mem.VirtualMemory()
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}

	d, _ := disk.Usage(path)

	if nodeType == "mn" {
		if (float64(v.Total)/float64(GB)) > float64(1.5) && (float64(d.Total)/float64(GB)) > float64(38.00) {
			log.Info("ethoFS - resource requirements met", "node type", nodeType)
		} else {

			errorMessage := ""

			if (float64(v.Total) / float64(GB)) < float64(1.5) {
				errorMessage = "not enough memory"
			} else if (float64(d.Total) / float64(GB)) < float64(38.00) {
				errorMessage = "not enough storage space"
			}
			log.Error("ethoFS - resource requirements not met - exiting", "node type", nodeType, "error", errorMessage)
			os.Exit(0)
		}
	} else if nodeType == "sn" {
		if (float64(v.Total)/float64(GB)) > float64(0.75) && (float64(d.Total)/float64(GB)) > float64(18.00) {
			log.Info("ethoFS - resource requirements met", "node type", nodeType)
		} else {

			errorMessage := ""

			if (float64(v.Total) / float64(GB)) < float64(0.75) {
				errorMessage = "not enough memory"
			} else if (float64(d.Total) / float64(GB)) < float64(18.00) {
				errorMessage = "not enough storage space"
			}
			log.Error("ethoFS - resource requirements not met - exiting", "node type", nodeType, "error", errorMessage)
			os.Exit(0)
		}
	} else if nodeType == "gn" {
		if (float64(v.Total)/float64(GB)) > float64(3.0) && (float64(d.Total)/float64(GB)) > float64(70.00) {
			log.Info("ethoFS - resource requirements met", "node type", nodeType)
		} else {

			errorMessage := ""

			if (float64(v.Total) / float64(GB)) < float64(3.0) {
				errorMessage = "not enough memory"
			} else if (float64(d.Total) / float64(GB)) < float64(70.00) {
				errorMessage = "not enough storage space"
			}
			log.Error("ethoFS - resource requirements not met - exiting", "node type", nodeType, "error", errorMessage)
			os.Exit(0)
		}
	}
}
