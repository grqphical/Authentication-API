package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CreateAccountResult struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

var accessToken string

func TestCreateAccountRoute(t *testing.T) {
	os.Setenv("SECRET", "78392078hnvq5nh6c13m4t0y8m4q")
	router := CreateRouter()

	data := url.Values{}
	data.Set("username", "JohnDoe")
	data.Set("password", "Password123!")
	data.Set("email", "johndoe@example.com")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/authentication/create-account", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, 202, w.Code)

	var result CreateAccountResult

	err := json.NewDecoder(w.Result().Body).Decode(&result)
	HandleError(err)

	accessToken = result.AccessToken
}

func TestLoginRoute(t *testing.T) {
	router := CreateRouter()

	data := url.Values{}
	data.Set("username", "JohnDoe")
	data.Set("password", "Password123!")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/authentication/", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, 202, w.Code)
}

func TestAuthenticationRoute(t *testing.T) {
	router := CreateRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Token", accessToken)

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	database.Clear()
}
