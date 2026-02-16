package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func getConnections() {
	cmd := exec.Command("ss", "-tnp")

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	scanner.Scan()

	fmt.Println("Remote IP Addresses from active/listening sockets:")
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}
		peer := fields[5]
		if peer == "0.0.0.0:*" || peer == "[::]:*" || peer == "*.*" {
			continue
		}

		ipPort := strings.Split(peer, ":")
		if len(ipPort) >= 2 {
			ip := ipPort[0]
			ip = strings.Trim(ip, "[]")
			fmt.Printf(" %s (from line: %s)\n", ip, line)
		}
	}
}
