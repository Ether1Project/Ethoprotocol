package main

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// Variable declaration for node storage status/communication
var MaxStorage = int(0)
var UsedStorage = int(0)
var selfMaxStorage = int(0)
var selfUsedStorage = int(0)
