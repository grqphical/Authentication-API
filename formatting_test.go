package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorToString(t *testing.T) {
	const errMessage string = "Example error"
    err := errors.New(errMessage)

    message := ErrorToString(err)
    assert.Equal(t, message, errMessage)
}

func TestGenerateHTTPError(t *testing.T) {
    const errMessage string = "Example Error"
    var expectedValue = map[string]string{"status" : "400", "message" : errMessage}

    jsonError := GenerateHTTPError(400, errMessage)
    assert.Equal(t, expectedValue, jsonError)
}
