package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "123456as45"
	for _, c := range s {
		if unicode.IsNumber(c) == false {
			fmt.Println("not number")
		}
	}
}
