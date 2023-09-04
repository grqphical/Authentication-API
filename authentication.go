package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Used to check if the incoming request has an authentication token and if it's valid
//
// If it is valid, it will continue the request normally, otherwise it will return either a 400 for a request missing a token or a 401 if the token
// is not valid
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Token") == "" {
			c.String(http.StatusBadRequest, "No token found in header. Please add Token to your header in order to be authenticated")
			c.Abort()
			return
		}

		_, err := ValidateAccessToken(c.Request.Header.Get("Token"))

        if err == errors.New(TOKEN_EXPIRED_ERR) {
            c.Redirect(http.StatusTemporaryRedirect, "/refresh-token")
        }
        if err != nil {
            c.String(http.StatusUnauthorized, ErrorToString(err))
            c.Abort()
            return
        }

        c.Next()
	}
}

// Takes in a password and generates a hash from it
//
// If it hashes successfully, it returns the hash and nil.
// Otherwise it returns a blank string and the error
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return "", errors.New("failed to hash password")
	}

	hash := string(bytes)

	return hash, nil
}

// Verifies if a password and hash match
//
// If they match return true, otherwise return false
func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
