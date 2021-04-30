package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string `json:"firstName"`
}

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

var key []byte

func Valid(u *UserClaims) error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("time has expired")
	}
	if u.SessionID == 0 {
		return fmt.Errorf("session is invalid")
	}
	return nil
}

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
	fmt.Printf("password hash of %s is %s\n", passwordTest, passwordHash)

	err = comparePassword(passwordTest, passwordHash)
	if err != nil {
		log.Fatalln("password does not match")
	}
	fmt.Println("Great, passwords match")

	signMessage([]byte(passwordTest))
	fmt.Println("Great, signed password ")
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

func createToken(c *UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, _ := token.SignedString(key)
	return signedToken, nil
}

func parseToken(signedToken string) (*UserClaims, error) {
	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(signedToken, claims,
		func (token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() == jwt.SigningMethodHS512.Alg() {
				return nil, fmt.Errorf("badness in signing alg")
			}
			return key, nil
		})
	if err != nil {
		return nil, fmt.Errorf("error in parseToken :%w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("error in parse token %w", err)
	}

	return token.Claims.(*UserClaims), nil

}