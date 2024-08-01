package net

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

/*
  Description: test client / server using TCP
*/

func TestTCPconn(t *testing.T) {
	go tcpServer(t, "5545")
	b := make([]string, 0)
	b = append(b, "tcpline1")
	b = append(b, "tcpline2")
	b = append(b, "STOP")
	tcpClient(t, "127.0.0.1:5545", b)
	t.Log("done with TestTCPconn")
}

func tcpClient(t *testing.T, connect string, lines []string) {
	c, err := net.Dial("tcp", connect)
	if err != nil {
		t.Error("client:", err)
		return
	}

	t.Logf("client: The TCP server is %s\n", c.RemoteAddr().String())

	for _, line := range lines {
		fmt.Fprintf(c, line+"\n") // write to "c"
		if strings.TrimSpace(line) == "STOP" {
			t.Log("TCP client exiting...")
			time.Sleep(1 * time.Second)
			return
		}
		reply, _ := bufio.NewReader(c).ReadString('\n')
		t.Logf("client: reply %s", reply)
	}
}

func tcpServer(t *testing.T, port string) {

	l, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(l net.Listener) {
		_ = l.Close()
	}(l)

	c, err := l.Accept()
	if err != nil {
		t.Error("server:", err)
		return
	}

	for {
		buffer, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			t.Error("server:", err)
			return
		}
		message := buffer[0 : len(buffer)-1]
		t.Log("server -> ", message)
		if strings.TrimSpace(message) == "STOP" {
			t.Log("Exiting TCP server!")
			return
		}

		t := time.Now()
		data := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(data))
	}
}
