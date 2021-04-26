package main

import (
	//"encoding/json"
	//"fmt"
	//"go/ast"
	//"log"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string `json:"firstName"`
}

func main() {
	//p1 := person{
	//	First: "Matthew",
	//}
	//
	//fmt.Println("Person one name ", p1.First)
	//
	//p2 := person{
	//	First: "Mark",
	//}
	//
	//persons := []person{p1, p2}
	//
	//personsJson, err := json.Marshal(persons)
	//if err != nil {
	//	log.Panic(err) //panic on a programmer error
	//}
	//fmt.Println("JSON of person ... ", string(personsJson))
	//
	//persons2 := []person{}
	//
	//err = json.Unmarshal(personsJson, &persons2)
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Println("Converted back to person struct ", persons2)

	http.HandleFunc("/encode", encodeHandler)
	http.HandleFunc("/decode", decodeHandler)
	http.ListenAndServe(":9000", nil)
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {

	p1 := person{
		First: "Matthew",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoding was bad", err)
	}

}

func decodeHandler(w http.ResponseWriter, r *http.Request) {

	p1 := person{}
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		return
	}
	fmt.Println("Decode object", p1)


}
