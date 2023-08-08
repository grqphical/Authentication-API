package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Used to represent an account in the project
type Account struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
	Email        string `json:"email"`
}

var accounts []Account = []Account{}
var secret []byte

// Basic API route to test the authentication with
func Index(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// Handles account creation
func CreateAccount(c *gin.Context) {

	c.Request.ParseMultipartForm(1000)
	username := c.Request.Form.Get("username")
	password := c.Request.Form.Get("password")
	email := c.Request.Form.Get("email")

	err := ValidateUsername(username)

	if err != nil {
		c.JSON(http.StatusBadRequest, GenerateHTTPError(http.StatusBadRequest, ErrorToString(err)))
		return
	}

	err = ValidatePassword(password)

	if err != nil {
		c.JSON(http.StatusBadRequest, GenerateHTTPError(http.StatusBadRequest, ErrorToString(err)))
		return
	}

	hash, err := HashPassword(password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, GenerateHTTPError(http.StatusInternalServerError, ErrorToString(err)))
		return
	}

	err = ValidateEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, GenerateHTTPError(http.StatusBadRequest, ErrorToString(err)))
		return
	}

	accountObj := Account{Username: username, PasswordHash: hash, Email: email}
	accounts = append(accounts, accountObj)

	token, err := CreateToken(accountObj)

	if err != nil {
		c.JSON(http.StatusInternalServerError, GenerateHTTPError(http.StatusInternalServerError, ErrorToString(err)))
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"token": token})
}

// Handles logging into existing accounts
func Login(c *gin.Context) {
	c.Request.ParseMultipartForm(1000)
	username := c.Request.Form.Get("username")
	password := c.Request.Form.Get("password")

	for _, account := range accounts {
		if account.Username == username {
			if VerifyPassword(password, account.PasswordHash) {
				token, err := CreateToken(account)

				if err != nil {
					c.JSON(http.StatusInternalServerError, GenerateHTTPError(http.StatusInternalServerError, ErrorToString(err)))
					return
				}

				c.JSON(http.StatusAccepted, map[string]string{"token": token})
			}
		}
	}
}

func main() {
	LoadDotEnv()
	secret = []byte(os.Getenv("SECRET"))

	router := gin.Default()

	authenticationRequired := router.Group("/")

	authenticationRequired.Use(AuthenticationMiddleware())
	{
		authenticationRequired.GET("/", Index)
	}

	authenticationRoutes := router.Group("/authentication")

	authenticationRoutes.POST("/create-account", CreateAccount)
	authenticationRoutes.PUT("/authenticate", Login)

	router.Run("127.0.0.1:8000")
}
