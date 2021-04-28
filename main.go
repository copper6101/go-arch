package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string `json:"firstName"`
}
/**
curl -XGET http://localhost:9000/encode

curl -XPOST http://localhost:9000/decode --data '{"firstName":"Bob"}' -H "Content-Type:application/json"
curl -XPOST http://localhost:9000/decode --data '{"firstName":"Bob"}' -H "Content-Type:application/json" --user markm:password123

*/
func main() {

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

	auth, password, ok := r.BasicAuth()

	if ok {
		fmt.Printf("Basic Auth data is user %s & password %s", auth, password)
	} else {
		fmt.Println("No Basic Auth found in the request")
	}
}
