package main

import (
	"flag"
	"fmt"
	"os"
)

const ESSID = "Casa_wifi"
const NET_INTERFACE = "wlan0"
const WPA_WORDLIST = "wordlist.txt"

const SSH_WORDLIST = "wordlist.txt"

var ROOT_PASS string

func init() {
	ROOT_PASS = os.Getenv("ROOT_PASSWORD")
}

func main() {
	wpa := flag.Bool("wpa", false, "execute automated wpa attack")
	ssh := flag.Bool("ssh", false, "execute automated ssh attack")
	flag.Parse()

	if *wpa {
		wpaAttack()
	} else if *ssh {
		sshAttack()
	} else {
		fmt.Println("theres no default attack, run this program with -ssh or -wpa flag")
	}
}
