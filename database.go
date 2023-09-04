package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Used to represent an account in the application
type Account struct {
    ID           int64
	Username     string 
	PasswordHash string 
	Email        string 
}

type Database struct {
    File string
    Connection *sql.DB
}

func NewDatabase(file string) Database {
    db, err := sql.Open("sqlite", file);
    HandleError(err)

    db.Exec("CREATE TABLE IF NOT EXISTS accounts (id INTEGER PRIMARY KEY, username STRING NOT NULL, password_hash STRING NOT NULL, email STRING NOT NULL)")

    return Database { File: file, Connection: db}
}

func (db Database) GetAccounts() ([]Account, error) {
    var accounts []Account

    rows, err := db.Connection.Query("SELECT * FROM accounts")
    HandleError(err)

    for rows.Next() {
        var acc Account
        if err = rows.Scan(&acc.ID, &acc.Username, &acc.PasswordHash, &acc.Email); err != nil {
            return nil, fmt.Errorf("Accounts: %v", err)
        }

        accounts = append(accounts, acc)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("Accounts: %v", err)
    }

    return accounts, nil
}

// Adds an account to the database and returns its id
func (db Database) AddAccount(account Account) (int64, error) {
    result, err := db.Connection.Exec("INSERT INTO accounts (username, password_hash, email) VALUES (?, ?, ?)", 
    account.Username, 
    account.PasswordHash,
    account.Email)


    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (db Database) GetAccountByID(id int64) (Account, error) {
    row := db.Connection.QueryRow("SELECT * FROM accounts WHERE id = ?", id)

    var acc Account

    if err := row.Scan(&acc.ID, &acc.Username, &acc.PasswordHash, &acc.Email); err != nil {
        return Account{}, err
    }

    return acc, nil
}

func (db Database) GetAccountByUsername(username string) (Account, error) {
    row := db.Connection.QueryRow("SELECT * FROM accounts WHERE username = ?", username)

    var acc Account

    if err := row.Scan(&acc.ID, &acc.Username, &acc.PasswordHash, &acc.Email); err != nil {
        return Account{}, err
    }

    return acc, nil
}

func (db Database) Clear() {
    _, err := db.Connection.Exec("DELETE FROM accounts")
    HandleError(err)
}
