package main

import (
	"fmt"
	"strconv"
)

func pinpan() {
	fmt.Println("PinPan")
	response := ""
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			response = "Pin Pan"
		} else if i%3 == 0 {
			response = "Pin"
		} else if i%5 == 0 {
			response = "Pan"
		} else {
			response = strconv.Itoa(i)
		}

		fmt.Println(response)
	}
}

func main() {
	pinpan()
}
