package main

import (
	"fmt"
	"time"
)

func pingPong(channel chan string) {
	for i := 0; ; i++ {
		if i % 2 == 0 {
			channel <- "Ping"
		} else {
			channel <- "Pong"
		}

		time.Sleep(850 * time.Millisecond)
	}
}

func print(channel chan string) {
	for {
		message := <- channel
		fmt.Println(message)
	}
}

func main() {
	channelPingPong := make(chan string)

	go pingPong(channelPingPong)
	go print(channelPingPong)

	select {}
}
