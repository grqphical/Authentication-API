package main

import (
	"encoding/json"
	"os"
)

// Used to represent an account in the project
type Account struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
	Email        string `json:"email"`
    UUID         string `json:"uuid"`
}

// Saves the accounts list to the disk
func SaveAccounts() {
	json, err := json.Marshal(accounts)
	HandleError(err)

	err = os.WriteFile("accounts.json", json, 0644)
	HandleError(err)
}

// Loads the accounts saved to the disk
func LoadAccounts() []Account {
	var accounts []Account

	data, err := os.ReadFile("accounts.json")
	if err != nil {
		err := os.WriteFile("accounts.json", []byte("[]"), 0644)
		HandleError(err)
		return []Account{}
	}

	err = json.Unmarshal(data, &accounts)
	HandleError(err)

	return accounts
}
