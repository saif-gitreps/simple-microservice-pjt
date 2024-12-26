package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["client"] = "John doe"
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString((MySigningKey))

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	token, err :=  GetJWT()
	fmt.Println(token)

	if err != nil {
		fmt.Fprintf(w, "Failed to generate, Error: %s", err.Error())
	} else {
		fmt.Fprintf(w, token)
	}
}

func handleRequest() {
	http.HandleFunc("/", Index) 

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}