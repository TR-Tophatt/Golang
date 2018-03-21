package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0 // 0.0 is important here (makes this a float 64)
	for _, v := range sf {
		fmt.Println(v)
		total += v
	}
	return total / float64(len(sf))
}
