package main

import "fmt"

func multiples(n int) {
	if n < 1 || n > 100 {
		fmt.Println("Apenas são aceitos números entre 1 e 100!")
	} else {
		fmt.Println("Múltiplos de ", n, "entre 1 e 100:")
		for i := 1; i <= 100; i++ {
			if i%n == 0 {
				fmt.Println(i)
			}
		}
	}
}

func main() {
	multiples(3)
}
