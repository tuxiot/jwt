package main

import (
	"fmt"
	"os"

	"github.com/tuxiot/jwt/pkg/jwt"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: jwt <secret key> <issuer>")
		os.Exit(0)
	}
	secret, issuer := os.Args[1], os.Args[2]

	token, err := jwt.GenerateToken([]byte(secret), issuer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(token)

}
