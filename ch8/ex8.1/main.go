package main

import (
	"fmt"
	"flag"
	"log"
	"net"
	"io"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return //e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "端口号 默认8000")
	flag.Parse()
	
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}