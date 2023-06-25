package main

import "fmt"

func main() {
	e := 1.0
	var lastVal float64 = 0
	for i := 1.0; i > 0; i++ {
		if lastVal == e {
			break
		} else {
			lastVal = e
		}

		var current_i float64 = i
		temp := 1.0
		for current_i > 0 {
			temp *= current_i
			current_i -= 1
		}

		e += 1 / temp
		fmt.Println(e)

	}
}
