package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsernameValidation(t *testing.T) {
    validUsername := "Username1"
    err := ValidateUsername(validUsername)

    assert.Equal(t, nil, err)

    invalidUsername := "ThisIsMoreThan24CharactersForAUsername"
    err = ValidateUsername(invalidUsername)
    
    assert.NotEqual(t, nil, err)
}

func TestValidEmail(t *testing.T) {
    validEmail := "john.doe@email.com"
    err := ValidateEmail(validEmail)
    
    assert.Equal(t, nil, err)

    invalidEmail := "john-doe@.com"
    err = ValidateEmail(invalidEmail)
    
    assert.NotEqual(t, nil, err)

}

func TestValidPassword(t *testing.T) {
    validPassword := "Password!123"
    err := ValidatePassword(validPassword)
    assert.Equal(t, nil, err)

    invalidPassword := "Password123"
    err = ValidatePassword(invalidPassword)
    assert.NotEqual(t, nil, err)
   
    invalidPassword = "Pas"
    err = ValidatePassword(invalidPassword)
    assert.NotEqual(t, nil, err)
    
    invalidPassword = "Password1238904-36ytpr;hwniy-08na9jwwtpwc8mgh7cjec0etx98o7a0w9hqw7c9y80cvhxq7whyc8j9o7ntkerihvnyo8cxjmqyktq097cmnwoyoqnc3x8"
    err = ValidatePassword(invalidPassword)
    assert.NotEqual(t, nil, err)

    invalidPassword = "passwords"
    err = ValidatePassword(invalidPassword)
    assert.NotEqual(t, nil, err)


}
