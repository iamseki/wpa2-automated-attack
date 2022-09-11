package main

import (
	"flag"
	"fmt"
	"os"
)

const ESSID = "Casa_wifi"
const NET_INTERFACE = "wlan0"
const WPA_WORDLIST = "wordlist.txt"

const SSH_WORDLIST = "wordlistssh.txt"
const IP_RANGE = "192.168.0.*"

var ROOT_PASS string

/*
	MANUAL STEPS:

	1. FIGURE OUT THE IP RANGE AND THE ESSID TO ATTACK
	2. GENERATE WORD LISTS
*/

func init() {
	ROOT_PASS = os.Getenv("ROOT_PASSWORD")
}

func main() {
	wpa := flag.Bool("wpa", false, "execute automated wpa attack")
	ssh := flag.Bool("ssh", false, "execute automated ssh attack")
	skipNmap := flag.Bool("skip-nmap", false, "skip everything and just execute hydra brute force into generated files")
	flag.Parse()

	if *wpa {
		wpaAttack()
	} else if *ssh {
		sshAttack(*skipNmap)
	} else {
		fmt.Println("theres no default attack, run this program with -ssh or -wpa flag")
	}
}
