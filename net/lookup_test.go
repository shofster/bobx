package net

import (
	"net"
	"testing"
)

func TestLookup(t *testing.T) {
	lookup(t, "google.com")
	lookup(t, "bobs-mini")
	lookup(t, "DellBob")
	lookup(t, "charon")
}

func lookup(t *testing.T, host string) {
	ips, err := net.LookupIP(host)
	if err != nil {
		t.Logf("Could not get IPs: %v\n", err)
	} else {
		for _, ip := range ips {
			t.Logf("%s IN A %s\n", host, ip.String())
		}
	}
}
