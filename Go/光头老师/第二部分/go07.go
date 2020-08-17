package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// 整数环绕
	var red uint8 = 255
	red++
	fmt.Println(red)
	var number int8 = 127
	number++
	fmt.Println(number)
	// 打印bit
	var green uint8 = 3
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)
	// 整数类型最大、小值 math包
	fmt.Println(math.MaxInt16)
	// 时间问题
	future := time.Unix(12622780800, 0)
	fmt.Println(future)
}
