// Package scanner executes the ss command and parses the results to provide the IP addresses for remote connections
package scanner

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/CodeZeroSugar/ip-guard/internal/process"
)

type IPScanner struct {
	runner process.Runner
}

func (i *IPScanner) GetRemoteIPs() (map[string][]string, error) {
	output, err := i.runner.Run("ss", "-tunp")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	scanner.Scan()

	fmt.Println("Remote IP Addresses from active/listening sockets:")
	hostMap := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}
		peer := fields[5]

		host, port, err := net.SplitHostPort(peer)
		if err != nil || port == "*" {
			continue
		}

		_, ok := hostMap[host]
		if !ok {
			hostMap[host] = []string{port}
		} else {
			hostMap[host] = append(hostMap[host], port)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return hostMap, nil
}
