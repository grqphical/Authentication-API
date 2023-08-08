package main

import "fmt"

// Used to convert an error to a string
func ErrorToString(err error) string {
	return fmt.Sprintf("%s", err)
}

// Generates the JSON for an HTTP error
func GenerateHTTPError(status int32, message string) map[string]string {
	return map[string]string{"status": string(status), "message": message}
}
