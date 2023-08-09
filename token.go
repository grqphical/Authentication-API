package main

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// Creates an access JWT token from an Account object
//
// Returns the token and nil on a successful encoding
// Otherwise it returns a blank string and an error
func CreateAccessToken(accountObj Account) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = accountObj.Username
	claims["hash"] = accountObj.PasswordHash
	claims["expiry"] = time.Now().Add(time.Minute * 15)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Takes in an access JWT, verifies the signature and returns the username associated with the account that is authenticated
//
// Returns either the username and nil or a blank string and an error if it failed to authenticate it
func ValidateAccessToken(tokenStr string) (string, error) {
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
			expirationTime, err := time.Parse(time.RFC3339, claims["expiry"].(string))
			if err != nil {
				return "", errors.New("could not validate token expiry")
			}

			if time.Now().After(expirationTime) {
				return "", errors.New(TOKEN_EXPIRED_ERR)
			}

			return claims["username"].(string), nil
		}

	}

	return "", errors.New("authentication error")
}

func CreateRefreshToken(accountObj Account) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = accountObj.UUID
	claims["expiry"] = time.Now().Add(time.Hour * 24 * 7)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateRefreshToken(tokenStr string) (string, error) {
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
			expirationTime, err := time.Parse(time.RFC3339, claims["expiry"].(string))
			if err != nil {
				return "", errors.New("could not validate token expiry")
			}

			if time.Now().After(expirationTime) {
				return "", errors.New("token is expired")
			}

			return claims["uuid"].(string), nil
		}

	}

	return "", errors.New("authentication error")
}
