package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	pid := os.Getpid()
	fmt.Printf("Listening for signals, pid: %d\n", pid)
	signal.Notify(c)
	s := <-c
	fmt.Printf("Recieved signal: %v\nExiting...\n", s)
}
