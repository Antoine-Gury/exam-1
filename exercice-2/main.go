package main

import "fmt"

func main() {
	fmt.Println(Countletters("Hello,World!"))
	fmt.Println(Countletters("123 456"))
	fmt.Println(Countletters("Golang1.21"))
}

func Countletters(s string) int {
	compter := 0
	for i := 0; i != len(s); i++ {
		if s[i] > 65 && s[i] < 122 {
			compter += 1
		}
	}
	return compter
}
