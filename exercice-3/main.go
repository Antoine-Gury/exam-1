package main

import "fmt"

func main() {
	decimal()
}
func decimal() {
	fmt.Println("0")
	chiffre := 0
	for i := 49; i <= 57; i++ {
		chiffre += 1
		fmt.Printf("%c", i)
		fmt.Println("\n")
	}

}
