package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Loads env variables in .env so they can be retrieved with os.Getenv()
func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const TOKEN_EXPIRED_ERR = "token is expired"

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

    	accessToken, err := CreateAccessToken(accountObj)

        if err != nil {
            c.JSON(http.StatusInternalServerError, GenerateHTTPError(http.StatusInternalServerError, ErrorToString(err)))
            return
        }

        refreshToken, err := CreateRefreshToken(accountObj)

        if err != nil {
            c.JSON(http.StatusInternalServerError, GenerateHTTPError(http.StatusInternalServerError, ErrorToString(err)))
            return
        }

        c.JSON(http.StatusAccepted, map[string]string{"refreshToken": refreshToken, "accessToken" : accessToken})
        return


	SaveAccounts()

}

// Handles logging into existing accounts
func Login(c *gin.Context) {
	c.Request.ParseMultipartForm(1000)
	username := c.Request.Form.Get("username")
	password := c.Request.Form.Get("password")

	for _, account := range accounts {
		if account.Username == username {
			if VerifyPassword(password, account.PasswordHash) {
				accessToken, err := CreateAccessToken(account)

				if err != nil {
					c.JSON(http.StatusInternalServerError, GenerateHTTPError(http.StatusInternalServerError, ErrorToString(err)))
					return
				}

				refreshToken, err := CreateRefreshToken(account)

				if err != nil {
					c.JSON(http.StatusInternalServerError, GenerateHTTPError(http.StatusInternalServerError, ErrorToString(err)))
					return
				}

				c.JSON(http.StatusAccepted, map[string]string{"refreshToken": refreshToken, "accessToken" : accessToken})
				return
			}
		}
	}

	c.JSON(http.StatusNotFound, GenerateHTTPError(http.StatusNotFound, "User doesn't not exist"))
}

// API Route to get a new access token if the current one is expired
func RefreshAccessToken(c *gin.Context){
    c.Request.ParseMultipartForm(1000)
    refreshToken := c.Request.Form.Get("refreshToken")
   
    uuid, err := ValidateRefreshToken(refreshToken)

    if err != nil {
        c.JSON(http.StatusBadRequest, GenerateHTTPError(http.StatusBadRequest, ErrorToString(err)))
        return
    }
    
    for _, account := range accounts {
        if account.UUID == uuid {
            token, err := CreateAccessToken(account)

            if err != nil {
                c.JSON(http.StatusBadRequest, GenerateHTTPError(http.StatusBadRequest, ErrorToString(err)))
                return
            }

            c.JSON(http.StatusAccepted, map[string]string{"accessToken" : token})
            return
        }
    }
}

func main() {
	LoadDotEnv()
	secret = []byte(os.Getenv("SECRET"))

	accounts = LoadAccounts()

	router := gin.Default()

	authenticationRequired := router.Group("/")

	authenticationRequired.Use(AuthenticationMiddleware())
	{
		authenticationRequired.GET("/", Index)
	}

	authenticationRoutes := router.Group("/authentication")

	authenticationRoutes.POST("/create-account", CreateAccount)
    authenticationRoutes.POST("/refresh-token", RefreshAccessToken)
	authenticationRoutes.PUT("/", Login)

	router.Run("127.0.0.1:8000")
}
