package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	SignKey   []byte
	publicKey []byte
)

type Token struct {
	Token string `json:"token"`
}

func init() {
	var err error
	SignKey, err = ioutil.ReadFile("app.rsa")
	if err != nil {
		log.Print(err)
	}
	publicKey, err = ioutil.ReadFile("app.rsa.pub")
	if err != nil {
		log.Print(err)
	}
	//log.Println("Public : ", signKey, "Private :", publicKey)
}

func returnInterface(data []byte) interface{} {
	return data
}

func main() {
	//log.Println(string(signKey), string(publicKey))
	tokenData := jwt.New(jwt.SigningMethodRS512)
	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenData.Claims = claims
	claims["StudentInfo"] = struct {
		Name string
		Role string
	}{"Ravi", "Student"}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(SignKey)
	if err != nil {
		log.Println(err)
	}
	tokenString, err := tokenData.SignedString(signKey)
	if nil != err {
		log.Println("Error while signing the token")
		log.Printf("Error signing token: %v\n", err)

	}
	response := Token{tokenString}
	log.Println(reflect.TypeOf(tokenString))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return signKey, nil
	})

	if err != nil {
		log.Println("Working FIne")
	} else {
		fmt.Println("Not working well")
	}
	log.Println(reflect.TypeOf(response), token)
}
