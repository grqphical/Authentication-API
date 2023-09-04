package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccessToken(t *testing.T) {
    account := Account{ID: 0, Username: "JohnDoe", PasswordHash: "6397uhpdasf0nc64780qcn", Email: "john.doe@example.com"}
    _, err := CreateAccessToken(account)

    assert.Equal(t, err, nil)
}

func TestValidateAccessToken(t *testing.T) {
    account := Account{ID: 0, Username: "JohnDoe", PasswordHash: "6397uhpdasf0nc64780qcn", Email: "john.doe@example.com"}
    token, err := CreateAccessToken(account)
    assert.Equal(t, err, nil)

    username, err := ValidateAccessToken(token)
    assert.Equal(t, err, nil)
    assert.Equal(t, username, account.Username)

    notTheToken := "j8963980pwehf8q9nby8569420-^@0nua"
    _, err = ValidateAccessToken(notTheToken)
    assert.NotEqual(t, err, nil)
}

func TestCreateRefreshToken(t *testing.T) {
    account := Account{ID: 0, Username: "JohnDoe", PasswordHash: "6397uhpdasf0nc64780qcn", Email: "john.doe@example.com"}
    _, err := CreateRefreshToken(account)
    assert.Equal(t, err, nil)
}

func TestValidateRefreshToken(t *testing.T) {
    account := Account{ID: 0, Username: "JohnDoe", PasswordHash: "6397uhpdasf0nc64780qcn", Email: "john.doe@example.com"}
    token, err := CreateRefreshToken(account)
    assert.Equal(t, err, nil)

    id, err := ValidateRefreshToken(token)
    assert.Equal(t, err, nil)
    assert.Equal(t, id, account.ID)

    notTheToken := "7623899pudhiqhnf86"
    _, err = ValidateRefreshToken(notTheToken)
    assert.NotEqual(t, err, nil)
}
