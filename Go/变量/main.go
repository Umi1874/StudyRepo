package main

import "fmt"

func main() {
	var a = "Runoob"
	fmt.Println(a)
	var b, c = 1, 2
	fmt.Println(b, c)
	d := 33
	fmt.Println(d)
	_, numb, strs := numbers()
	fmt.Println(numb, strs)
}

func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
