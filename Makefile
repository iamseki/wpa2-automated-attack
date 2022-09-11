attack-wpa:
	go run *.go --wpa

attack-ssh:
	go run *.go --ssh

attack-ssh-skip-nmap:
	go run *.go --ssh --skip-nmap