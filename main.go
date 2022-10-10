package main

import "calculator/calc"

func main() {
	str := "3(2+1)-9*2^3"
	c := calc.NewCalculator(str)
	println(c.Calculate())
}
