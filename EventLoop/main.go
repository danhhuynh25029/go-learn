package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	a := make(chan int, 3)
	sigChan := make(chan os.Signal, 1)
	go EventLoop(a)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
}

func EventLoop(ch chan int) {
	for {
		time.Sleep(time.Second)
		if len(ch) > 0 {
			fmt.Println("message : ", ch)
		} else {
			fmt.Println("do not message")
		}
	}
}
