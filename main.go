package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	toString := base64.StdEncoding.EncodeToString([]byte("usermark:password123"))
	fmt.Println(toString)

	_, err := base64.StdEncoding.DecodeString(toString)
	if err != nil {
		fmt.Errorf("This could not be decoded %s %w ", toString, err)
	}

}