package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string  `json:"firstName"`
}

func main() {
	p1 := person{
		First: "Matthew",
	}

	fmt.Println("Person one name ", p1.First)

	p2 := person{
		First: "Mark",
	}

	persons := []person{p1, p2}

	personsJson, err := json.Marshal(persons)
	if err != nil {
		log.Panic(err)  //panic on a programmer error
	}
	fmt.Println("JSON of person ... ", string(personsJson))

	persons2 := []person{}

	err = json.Unmarshal(personsJson, &persons2)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Converted back to person struct ", persons2)


}
