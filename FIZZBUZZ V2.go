package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch i {
		case i%3 == 0 && i%5 == 0:
			fmt.Println(i, "FIZZBUZZ")
		case i%5 == 0:
			fmt.Println(i, "BUZZ")
		case i%3 == 0:
			fmt.Println(i, "FIZZ")
		default:
			fmt.Println(i)
		}
	}

}
