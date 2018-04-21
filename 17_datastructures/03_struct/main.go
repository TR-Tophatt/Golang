package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//Person : a person
type Person struct {
	First       string
	Last        string
	Age         int `json:"wisdom score"`
	notExported int
}

func (p Person) fullName() string {
	return p.First + p.Last
}

//func (receiver) funcName(parameters) return {body}
// receiver attaches a function to a type; any value of type receiver has access

func main() {
	p1 := Person{
		First:       "James",
		Last:        "Bond",
		Age:         20,
		notExported: 007,
	}
	bs, ok := json.Marshal(p1)
	fmt.Println("Marshalling")
	fmt.Println(bs)
	fmt.Println(ok)
	fmt.Println(string(bs))
	fmt.Println("----------")

	var p2 Person
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)

	bs2 := []byte(`{"First":"Tater", "Last":"Rice", "wisdom score":"20"}`)

	json.Unmarshal(bs2, &p2)
	fmt.Println("Unmarshalling")
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)
	fmt.Printf("%T \n", p2)

	fmt.Println("----------")
	fmt.Println("Encoding")

	json.NewEncoder(os.Stdout).Encode(p1)

	fmt.Println("----------")
	fmt.Println("decoding")

	var p3 Person
	rdr := strings.NewReader(`{"First":"Jessica", "Last":"Rice", "wisdom score":"25"}`)
	json.NewDecoder(rdr).Decode(&p3)

	fmt.Println(p3.First)
	fmt.Println(p3.Last)
	fmt.Println(p3.Age)
}

//godoc.org/encoding/json
