package main

import (
	"net/textproto"
	"fmt"
)

func main() {
	s := textproto.TrimString(" xxx  ")
	fmt.Println(s)

	b := []byte(" xxxx xxx")
	bb := textproto.TrimBytes(b)
	fmt.Println(string(bb))

	conn, err := textproto.Dial("115.239.210.27", ":443")
	if err != nil {

	}
	fmt.Printf("info is (%+v)\n",conn)
	st, err := conn.ReadLine()
	if err != nil {

	}
	fmt.Println(st)
}
