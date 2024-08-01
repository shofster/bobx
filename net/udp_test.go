package net

import (
	"math/rand"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"
)

/*
  Description: test client / server using UDP
*/

func TestUDPconn(t *testing.T) {
	go udpServer(t, "5545")
	b := make([]string, 0)
	b = append(b, "udpline1")
	b = append(b, "udpline2")
	b = append(b, "STOP")
	udpClient(t, "127.0.0.1:5545", b)
}

func udpClient(t *testing.T, connect string, lines []string) {
	s, err := net.ResolveUDPAddr("udp4", connect)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		t.Error("client:", err)
		return
	}

	t.Logf("client: The UDP server is %s\n", c.RemoteAddr().String())
	defer func(c *net.UDPConn) {
		_ = c.Close()
	}(c)

	for _, line := range lines {
		data := []byte(line + "\n")
		_, err = c.Write(data)
		if strings.TrimSpace(string(data)) == "STOP" {
			t.Log("Exiting UDP client!...")
			time.Sleep(1 * time.Second)
			return
		}
		time.Sleep(2 * time.Second)

		if err != nil {
			t.Error("client:", err)
			return
		}
		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			t.Error("client:", err)
			return
		}
		t.Logf("client: reply %s\n", string(buffer[0:n]))
	}
}

func udpServer(t *testing.T, port string) {

	var random = func(min, max int) int {
		return rand.Intn(max-min) + min
	}

	s, err := net.ResolveUDPAddr("udp4", ":"+port)
	if err != nil {
		t.Error("server:", err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		t.Error("server:", err)
		return
	}

	defer func() {
		_ = connection.Close()
	}()

	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		t.Log("server -> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			t.Log("UDP client exiting...")
			return
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		t.Logf("  data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			t.Error("server:", err)
			return
		}
	}
}
