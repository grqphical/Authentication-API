package main

import "log"

const TOKEN_EXPIRED_ERR = "token is expired"

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
