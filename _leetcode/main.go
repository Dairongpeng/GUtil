package main

import "fmt"

func main() {

	s := "anaksdjdndlowndm"
	bytes := []byte(s)

	for i, v := range bytes {
		fmt.Printf("%d : %c", i, v)
	}
}
