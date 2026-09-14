package main

import (
	"fmt"
	"time"
)

// Задание 5.

func sumdiff(a float64, b float64) (float64, float64) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// Задание 6.
func avgg(a float64, b float64, c float64) float64 {
	avg := (a + b + c) / 3
	return avg
}

func main() {
	fmt.Println("Задание 1:", time.Now())

	var a int = 64
	var b float64 = 5.5
	var c string = "blaBlabla"
	var d bool = true
	fmt.Println("Задание 2:", a, b, c, d)

	a1 := 64
	b2 := 5.5
	c3 := "blaBlabla"
	d4 := true
	fmt.Println("Задание 3:", a1, b2, c3, d4)

	a11 := 6
	b11 := 2
	c11 := a11 + b11
	d11 := a11 - b11
	e := a11 * b11
	f := a11 / b11
	g := a11 % b11
	fmt.Println("Задание 4:", c11, d11, e, f, g)

	sum, diff := sumdiff(10.1, 1.5)
	fmt.Println("Задание 5:", sum, diff)

	avg := avgg(3, 4, 6)
	fmt.Println("Задание 6:", avg)
}
