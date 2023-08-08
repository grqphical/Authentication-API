package main

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

// Creates a JWT from an Account object
//
// Returns the token and nil on a successful encoding
// Otherwise it returns a blank string and an error
func CreateToken(accountObj Account) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = accountObj.Username
	claims["hash"] = accountObj.PasswordHash
	claims["authorized"] = true

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Takes in a JWT, verifies the signature and returns the username associated with the account that is authenticated
//
// Returns either the username and nil or a blank string and an error if it failed to authenticate it
func ValidateJWT(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("failed to authenticate token signature")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)

		if ok {
			return claims["username"].(string), nil
		}

	}

	return "", errors.New("authentication error")
}
