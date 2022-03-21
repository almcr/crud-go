package models

import (
	"log"

	"github.com/almcr/crud-go/helper"
)

type TokenPair struct {
	Token        string
	RefreshToken string `json:"refresh_token"`
}

// To track authorized users
type Accounts map[string]string
type Tokens map[string]TokenPair

var AuthUsers Accounts
var AuthTokens Tokens

func SetDefaultUsers() {
	AuthUsers = Accounts{"user0": "password"}
	AuthTokens = Tokens{}
	t, r, err := helper.GenerateTokens("user0")
	if err != nil {
		log.Fatal(err)
		return
	}

	AuthTokens["user0"] = TokenPair{t, r}
}
