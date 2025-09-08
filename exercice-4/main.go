package main

import "fmt"

func main() {
	fmt.Println(AffichagePGCD(24, 32))
}

func AffichagePGCD(a int, b int) int {
	PGCD := 1
	if a < b {
		for i := 9; i > 0; i-- {
			if a%i == 0 && b%i == 0 {
				PGCD = i
				fmt.Printf("%c", i)
			}
		}

	}
	return
}
