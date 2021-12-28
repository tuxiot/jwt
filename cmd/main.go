package main

import (
	"fmt"
	"os"

	"github.com/tuxiot/jwt/pkg/jwt"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: jwt <secret key> <issuer> <username>")
		os.Exit(0)
	}
	secret, issuer, username := os.Args[1], os.Args[2], os.Args[3]

	token, err := jwt.GenerateToken([]byte(secret), issuer, username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(token)

}
