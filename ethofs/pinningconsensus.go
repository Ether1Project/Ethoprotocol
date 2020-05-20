package main

import (
	"fmt"
	"strconv"
)

type NodePins struct {
	NodeID      string
	Pins        []string
	LatestBlock int
}

func (n NodePins) PrintNodePins() {
	numberOfPins := strconv.Itoa(len(n.Pins))
	fmt.Println("Locally Pinned Content (Pin Count): " + numberOfPins)
}
