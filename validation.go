package main

import (
	"errors"
	"fmt"
	"regexp"
)

// Ensures that the given username matches the username requirements
//
// All usernames must be shorter than 24 characters
// Returns nil if the username is valid and returns an error if it isnt'
func ValidateUsername(username string) error {
	const USERNAME_MAX_LENGTH int = 24

	if len(username) >= USERNAME_MAX_LENGTH {
		return fmt.Errorf("username must be less than %d characters", USERNAME_MAX_LENGTH)
	}

	return nil
}

// Ensures passwords are valid meaning they are in between 8 and 70 characters,
// have at least one capital, one number and one special character
// Example valid password: Password123!
//
// Returns nil if the username is valid and returns an error if it isnt'
func ValidatePassword(password string) error {
	const PASSWORD_MIN_LENGTH = 8
	const PASSWORD_MAX_LENGTH = 70

	// Make sure password is within accepted lengths
	if len(password) > PASSWORD_MAX_LENGTH {
		return errors.New("password too long. Maximum of 70 characters")
	} else if len(password) < PASSWORD_MIN_LENGTH {
		return errors.New("password too short. Minimum of 8 characters")
	}

	hasNumber, hasCapital, hasSpecialChar := false, false, false

	// Check to see if they used one capital in the password
	for _, char := range password {
		if 'A' <= char && char <= 'Z' {
			hasCapital = true
		}
	}

	// Check to see if they used one number in the password
	for _, char := range password {
		if '0' <= char && char <= '9' {
			hasNumber = true
		}
	}

	// Check to see if they used one special character in the password
	regex := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)

	if regex.MatchString(password) {
		hasSpecialChar = true
	}

	if hasNumber && hasCapital && hasSpecialChar {
		return nil
	} else {
		return errors.New("password must contain one number, one capital and one special character")
	}
}

func ValidateEmail(email string) error {
    // This regex determines if an email has valid characters and is in the correct format: name@domain.tld
	regex := regexp.MustCompile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")

	if regex.MatchString(email) {
		return nil
	}
	return errors.New("invalid email address")
}
