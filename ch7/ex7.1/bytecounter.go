package main

import (
	"fmt"
	"bufio"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p))
    return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	sum := 0
	for scanner.Scan() {
		sum++
	}

	*w += WordCounter(sum)
	return sum,nil
}


type LineCounter int
func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	sum := 0
	for scanner.Scan() {
		sum++
	}
	*l += LineCounter(sum)
	return sum, nil
}

func main() {
	var l LineCounter
	p := []byte("one\ntwo\nthree\n")
	n, _ := l.Write(p)
	fmt.Println(n)

	var w WordCounter
	data := []byte("The upcoming word is sp")
	n, _ = w.Write(data)
	fmt.Println(n)
}