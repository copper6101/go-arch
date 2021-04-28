package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string `json:"firstName"`
}

func main() {
	var secret string
	secret = base64.StdEncoding.EncodeToString([]byte("userName:password123"))
	fmt.Println(secret)

	passwordTest := "secret007"

	passwordHash, err := hashPassword(passwordTest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("password hash of %s is %s\n", passwordTest, passwordHash)

	err = comparePassword(passwordTest, passwordHash)
	if err != nil {
		log.Fatalln("password does not match")
	}
	fmt.Println("Great, passwords match")
}

func hashPassword(password string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error generating hash from password: %w", err)
	}
	return result, nil
}

func comparePassword(password string, hashPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPassword, []byte(password))
	if err != nil {
		return fmt.Errorf("error comparing password: %w", err)
	}
	return nil
}