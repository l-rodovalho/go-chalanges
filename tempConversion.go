package main

import "fmt"

func main() {
	celsius := 100

	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273

	fmt.Println("Temperatura de ebulição da água em Celsius:", celsius)
	fmt.Println("Temperatura de ebulição da água em Fahrenheit:", fahrenheit)
	fmt.Println("Temperatura de ebulição da água em Kelvin:", kelvin)
}
