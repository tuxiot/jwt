package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

const API_CLIENT = "apiclient"

// GenerateToken generates a new JWT token
func GenerateToken(secrect []byte, issuer string) (token string, err error) {
	// Create the Claims
	claims := &jwt.StandardClaims{
		Issuer: issuer,
	}

	tokenWithClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := tokenWithClaim.SignedString(secrect)
	return ss, err
}

// VerifyToken verifies the JWT token with given secret key. HMAC only
func VerifyToken(tokenString string, secretKey []byte) (err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return
}
