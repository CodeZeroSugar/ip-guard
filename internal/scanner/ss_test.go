package scanner

import (
	"testing"

	"github.com/CodeZeroSugar/ip-guard/internal/process"
)

func TestGetRemoteIPs(t *testing.T) {
	i := IPScanner{Runner: process.OSRunner{}}

	results, err := i.GetRemoteIPs()
	if err != nil {
		t.Log(err)
	}
	for host, ports := range results {
		t.Logf("%s : %v", host, ports)
	}
}
