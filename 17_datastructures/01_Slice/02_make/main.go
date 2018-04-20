package main

import "fmt"

func main() {
	records := make([][]string, 0)
	//student 1
	student1 := make([]string, 4)
	student1[0] = "Rice"
	student1[1] = "Taylor"
	student1[2] = "100.00"
	student1[3] = "74.00"
	student2 := make([]string, 4)
	student2[0] = "twede"
	student2[1] = "jessica"
	student2[2] = "100.00"
	student2[3] = "74.00"

	records = append(records, student1, student2)
	fmt.Println(records)

}
