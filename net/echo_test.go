package net

import (
	"io"
	"log"
	"net"
	"testing"
	"time"
)

/*
  Description: "echo" for 2 UDP sockets.
*/

func TestUDPecho(t *testing.T) {

	sender := &udpEchoer{}
	_ = sender.connect("127.0.0.1:5545", func(data []byte, err error) {
		t.Logf("sender %v got %s\n", err, string(data))
	})

	receiver := &udpEchoer{}
	_ = receiver.connect("127.0.0.1:5546", func(data []byte, err error) {
		t.Logf("receiver %v got %s\n", err, string(data))
	})

	_ = sender.send(receiver, []byte("message from sender to receiver"))
	time.Sleep(1 * time.Second)
	_ = receiver.send(sender, []byte("message from receiver to sender"))
	time.Sleep(1 * time.Second)
	if ping("udp", "127.0.0.1", "5546") == nil {
		t.Logf("%s %s %s\n", "127.0.0.1", "responding on port:", "5546")
	}
}

func ping(proto, host, port string) error {
	timeout := time.Duration(1 * time.Second)
	conn, err := net.DialTimeout(proto, host+":"+port, timeout)
	if err == nil {
		_ = conn.Close()
		return nil
	}
	//fmt.Printf("%s %s %s\n", host, "not responding", err.Error())
	return err
}

type udpEchoer struct {
	udpAddr  *net.UDPAddr
	udpConn  *net.UDPConn
	receiver func([]byte, error)
}

func (u *udpEchoer) addr() *net.UDPAddr {
	return u.udpAddr
}
func (u *udpEchoer) ip() []byte {
	return u.udpAddr.IP[12:]
}
func (u *udpEchoer) port() int {
	return u.udpAddr.Port
}
func (u *udpEchoer) connect(addr string, receiver func([]byte, error)) (err error) {
	u.receiver = receiver
	u.udpAddr, err = net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		return
	}
	if u.udpConn, err = net.ListenUDP("udp4", u.udpAddr); err == nil {
		// start listener function
		go func() {
			buffer := make([]byte, 1024)
			n, _, err := u.udpConn.ReadFromUDP(buffer)
			log.Println("receive:", string(buffer[0:n]))
			if u.receiver != nil {
				u.receiver(buffer[0:n], err)
			}
		}()
	}
	return
}
func (u *udpEchoer) send(receiver *udpEchoer, data []byte) error {
	log.Printf("send: %d bytes, from %v:%v  to %v:%v\n", len(data),
		u.ip(), u.port(), receiver.ip(), receiver.port())
	n, err := u.udpConn.WriteToUDP(data, receiver.addr())
	if n != len(data) {
		err = io.ErrShortWrite
	}
	return err
}
func (u *udpEchoer) close() {
	_ = u.udpConn.Close()
}
