package main

import "fmt"

var i string

func Printalphabet() {
	for i := "a"; i <= "z"; i++ {
		fmt.Println("%C")
		fmt.Println(i)
	}
}
