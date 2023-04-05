package main

import "fmt"

func main() {
	input := 14

	if input % 7 == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}