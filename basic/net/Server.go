package main

import (
	"net"
	"fmt"
)

var client_num int = 0

func main ()  {
	StartServer()
}
func StartServer() {

	l, err := net.Listen("tcp", ":1200")
	if err != nil {
		fmt.Println("err",err)
		return
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		client_num++
		fmt.Printf("A new Connection %d.\n", client_num)
		go handlerConnection(conn)
	}
}

func handlerConnection(conn net.Conn) {
	defer closeConnection(conn)

	readChannel := make(chan []byte, 1024)
	writeChannel := make(chan []byte, 1024)

	go readClientConnection(conn, readChannel)
	go writeServerConnection(conn, writeChannel)

	for {
		select {
		case data := <-readChannel:
			if string(data) == "bye" {
				return
			}
			writeChannel <- append([]byte("Back"), data...)
		}
	}
}

func writeServerConnection(conn net.Conn, channel chan []byte) {
	for {
		select {
		case data := <-channel:
			println("Write:", conn.RemoteAddr().String(), string(data))
			_, err := conn.Write(data)
			if err != nil {
				return
			}
		}
	}

}

func readClientConnection(conn net.Conn, channel chan []byte) {

	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			channel <- []byte("bye")
			break
		}
		println("Recei:", conn.RemoteAddr().String(), string(buffer[:n]))
		channel <- buffer[:n]
	}
}

func closeConnection(conn net.Conn) {
	conn.Close()
	client_num--
	fmt.Printf("Now, %d connections is alve.\n", client_num)
}
