package main

import "fmt"

func main() {
	/**
	A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array
	A slice type denotes the set of all slices of arrays of its element type
	The value of an uninitialized slice is nil
	Slices can change in size
	slices have length and capacity
	https://golang.org/ref/spec#Slice_types

	map
	key/value storage
	similar to "dictionary"
	unordered group of elements of one type called element type
	indexed by a set of unique keys of another type, key type
	uninitialized maps are nil
	https://golange.org/ref/spec#Map_types

	struct
	data structure, composite time, collects properties together
	https://golange.org/ref/spec#Struct_types

	**/

	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}
