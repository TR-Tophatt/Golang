package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...)
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
