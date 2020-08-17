package main

import (
	"fmt"
)

// string
func main() {
	fmt.Println("peace\naaa")
	fmt.Println(`peace\naaa`)
	fmt.Println(`----------------------------------------`)
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)
	fmt.Println(`----------------------------------------`)
	message := "shalom"
	for i := 0; i < len(message); i++ {
		fmt.Printf("%c \n", message[i])
	}
	fmt.Println(`----------------------------------------`)
	// 凯撒加密法 ROT13
	// range
	question := "aaddffgg"

	for _, c := range question {
		fmt.Printf("%v %c\n", c, c)
	}
}
