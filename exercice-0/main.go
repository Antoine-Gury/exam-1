package main

import "fmt"

func main() {
	printalphabet()
}
func printalphabet() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Println("\n")
}
