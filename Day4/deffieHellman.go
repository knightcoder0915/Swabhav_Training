package main

import (
	"fmt"
	"math"
)

func main() {
	p := 17
	q := 7
	aliceKey := 2
	x := int(math.Pow(float64(p), float64(aliceKey))) % q
	bobKey := 5
	y := int(math.Pow(float64(p), float64(bobKey))) % q

	ka := int(math.Pow(float64(y), float64(aliceKey))) % q
	fmt.Println("ka", ka)
	kb := int(math.Pow(float64(x), float64(bobKey))) % q
	fmt.Println("kb", kb)
}
