package main

import "fmt"

func main() {
	fmt.Println(AffichagePGCD(24, 32))
}

func AffichagePGCD(a int, b int) int {
	PGCD := 1
	for i := 1; i <= 9; i++ {
		if a/i == b/i {
			PGCD = i
			fmt.Printf("%c", i)
		}

	}
	return PGCD
}
