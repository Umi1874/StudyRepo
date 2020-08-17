package main

import (
	"fmt"
	"math/big"
)

// 非常大的数 big包
func main() {
	// big.Int
	lightSpeed := big.NewInt(299792)
	secondsPerday := big.NewInt(86400)
	fmt.Println(lightSpeed, secondsPerday)
	distance := new(big.Int)
	distance.SetString("24000000000000000000000000000000", 10)
	fmt.Println(distance)
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerday)
	fmt.Println(days)
	// big.Float
	// big.Rat
}
