package net

import (
	"fmt"
	"net"
	"testing"
	"time"
)

/*
  Description:  broadcast for UDP sockets
*/

const broadcastTestPort = ":8800"
const broadcastTestAddress = "192.168.1.255:8800"

func TestUDPbroadcast(t *testing.T) {
	go listener()
	meth1()
	// give time for close's to complete
	time.Sleep(300 * time.Millisecond)
	go listener()
	meth2()
	time.Sleep(300 * time.Millisecond)
	go listener()
	meth3()
	time.Sleep(300 * time.Millisecond)
	go listener()
	meth4()
}

func listener() {
	pc, err := net.ListenPacket("udp4", broadcastTestPort)
	if err != nil {
		fmt.Printf("listener ListenPacket error: %v\n", err)
		return
	}
	defer func(pc net.PacketConn) {
		err := pc.Close()
		if err != nil {
			fmt.Printf("** listener close error: %v\n", err)
		}
	}(pc)
	_ = pc.SetReadDeadline(time.Now().Add(2 * time.Second))

	buf := make([]byte, 1024)
	n, addr, err := pc.ReadFrom(buf)
	if err != nil {
		fmt.Printf("listener ReadFrom error: %v\n", err)
		return
	}
	fmt.Printf("from %s (%s) got: %s\n", addr.String(), addr.Network(), buf[:n])
}

func meth1() { // ListenPacket - preferred simple
	var err error

	remote, err := net.ResolveUDPAddr("udp4", broadcastTestAddress)
	if err != nil {
		panic(err)
	}

	initiator, err := net.ListenPacket("udp4", ":8828")
	if err != nil {
		panic(err)
	}
	defer func(initiator net.PacketConn) {
		_ = initiator.Close()
	}(initiator)

	_, err = initiator.WriteTo([]byte("meth1: data to transmit"), remote)
	if err != nil {
		panic(err)
	}
}

func meth2() { // ListenUDP
	var err error

	remote, err := net.ResolveUDPAddr("udp4", broadcastTestAddress)
	if err != nil {
		panic(err)
	}

	initiatorAddr, err := net.ResolveUDPAddr("udp4", ":8828")
	if err != nil {
		panic(err)
	}
	initiator, err := net.ListenUDP("udp4", initiatorAddr)
	if err != nil {
		panic(err)
	}
	defer func(initiator *net.UDPConn) {
		_ = initiator.Close()
	}(initiator)

	_, err = initiator.WriteTo([]byte("meth2: data to transmit"), remote)
	if err != nil {
		panic(err)
	}
}

func meth3() { // DialUDP
	var err error

	remote, err := net.ResolveUDPAddr("udp", broadcastTestAddress)
	if err != nil {
		panic(err)
	}

	initiatorAddr, err := net.ResolveUDPAddr("udp", ":8029")
	if err != nil {
		panic(err)
	}

	initiator, err := net.DialUDP("udp4", initiatorAddr, remote)
	if err != nil {
		panic(err)
	}
	defer func(initiator *net.UDPConn) {
		_ = initiator.Close()
	}(initiator)

	_, err = initiator.Write([]byte("meth3: data to transmit"))
	if err != nil {
		panic(err)
	}
}

func meth4() { // net.Dial() - most common - initiator port is syste generated
	var err error

	addr := fmt.Sprintf("127.0.0.1%s", broadcastTestPort)
	conn, err := net.Dial("udp", addr)
	if err != nil {
		panic(err)
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	_, err = conn.Write([]byte("meth4: data to transmit"))
	if err != nil {
		panic(err)
	}
}

/*

Unicast: from one source to one destination i.e. One-to-One
Broadcast: from one source to all possible destinations i.e. One-to-All
Multicast: from one source to multiple destinations stating an interest in
  receiving the traffic i.e. One-to-Many


In the UDP protocol, there is no distinction between clients and servers.
Creating a UDP socket or “connection” does not involve sending any packets.
A UDP client is simply the the initiator, the party that sends the first packet
 , rather than the responder, the party that receives the first packet.
The initiator necessarily knows the remote address a priori
 , since the initiator has to send the first packet.
The responder can learn the remote address when it receives the packet.

The common instantiation of a UDP client in Go is net.Dial("udp", address).
This returns a net.Conn object implemented by a net.UDPConn.
It provides both Read and Write methods.
*/

/*

https://github.com/aler9/howto-udp-broadcast-golang

All four methods work and their result is indistinguishable.
By looking at the Go source code, it is possible to assert that:

net.ListenPacket() and net.ListenUDP() use a nearly identical identical procedure,
 as they both call the same system functions,
 the only difference is that net.ListenPacket() converts the
 desired listen address (:8829) into an UDPAddr structure,
 while net.ListenUDP() requires directly an UDPAddr structure;

net.DialUDP() uses a different route and also provides a Write() function
 to write directly to the broadcast address.
 This could be confusing, as Go always work with WriteTo() and ReadFrom()
 when dealing with UDP connections.

Conclusion: I use with the ListenPacket() solution as it is the simpler one.

*/
