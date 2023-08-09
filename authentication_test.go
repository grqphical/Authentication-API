package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {
    password := "password123"
    wrongPassword := "password124"
	hash, err := HashPassword(password)
    HandleError(err)

    assert.Equal(t, VerifyPassword(password, hash), true)

    assert.Equal(t, VerifyPassword(wrongPassword, hash), false)
}
