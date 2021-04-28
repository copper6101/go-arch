package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string `json:"firstName"`
}

var key = []byte{}


func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	var secret string
	secret = base64.StdEncoding.EncodeToString([]byte("userName:password123"))
	fmt.Println(secret)

	passwordTest := "secret007"

	passwordHash, err := hashPassword(passwordTest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("password hash of %v is %v\n", passwordTest, passwordHash)

	err = comparePassword(passwordTest, passwordHash)
	if err != nil {
		log.Fatalln("password does not match")
	}
	fmt.Println("Great, passwords match")

	signedMessaged, _ := signMessage([]byte(passwordTest))

	fmt.Println(signedMessaged)

	same, _ := checkSignMessageIsSame(passwordTest, signedMessaged)

	fmt.Printf("Same signature %v\n", same)

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

func signMessage(message []byte) ([]byte, error) {

	hash := hmac.New(sha512.New, key)

	_, err := hash.Write(message)
	if err != nil {
		return nil, fmt.Errorf("error while hashing %w", err)
	}

	signature := hash.Sum(nil)
	return signature, nil
}

func checkSignMessageIsSame(message string, signature []byte) (bool, error) {

	newSig, _ := signMessage([]byte(message))
	same := hmac.Equal(newSig, signature)
	return same, nil
}