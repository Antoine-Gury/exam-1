package main

import "fmt"

func main() {
	PrintAlphabet()
}

func PrintAlphabet() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Println("\n")
}
