package main

import (
	"net"
	"os/signal"
	"os"
	"fmt"
	"syscall"
)

var (
	sigs = make(chan os.Signal, 1)
	done = make(chan bool, 1)
)

func  main()  {
	StartClient()
}
// StartClient
func StartClient() {

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:1200")
	if err != nil {
		<-done
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		<-done
	}
	writeChan := make(chan []byte, 1024)
	readChan := make(chan []byte, 1024)

	go readConnection(conn, readChan)
	go writeConnection(conn, writeChan)

	for {
		var s string
		fmt.Scan(&s)
		writeChan <- []byte(s)
	}
}

func readConnection(conn *net.TCPConn, channel chan []byte) {

	defer conn.Close()
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			<-done
			return
		}
		fmt.Println("Received from:", conn.RemoteAddr(), string(buffer[:n]))
	}
}

func writeConnection(conn *net.TCPConn, channel chan []byte) {

	defer conn.Close()
	for {
		select {
		case data := <-channel:
			_, err := conn.Write(data)
			if err != nil {
				<-done
				return
			}
			fmt.Println("Write to:", conn.RemoteAddr(), string(data))
		}
	}
}
