package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
		c.Close()
	}()
	line := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			line <- input.Text()
		}
	}()
	timeout := 10 * time.Second
	timer := time.NewTimer(timeout)

	for {
		select {
		case <-timer.C:
			fmt.Println("close conn")
			return
		case text := <-line:
			timer.Reset(timeout)
			wg.Add(1)
			go echo(c, text, 1*time.Second, wg)
		}
	}

}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
