package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "Matthew",
	}

	p2 := person{
		First: "Mark",
	}

	persons := []person{p1, p2}

	bs, err := json.Marshal(persons)
	if err != nil {
		log.Panic(err)  //panic on a programmer error
	}
	fmt.Println(string(bs))

}
