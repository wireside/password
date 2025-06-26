package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var availableLetterRunes = []rune(
	"abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*@!_.",
)

var (
	invalidUrlError   = errors.New("url is invalid")
	invalidLoginError = errors.New("login is empty")
)

type Account struct {
	login    string `json:"login" xml:"test"`
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.password)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = availableLetterRunes[rand.IntN(len(availableLetterRunes)-1)]
	}

	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (
	*AccountWithTimeStamp, error,
) {
	if login == "" {
		return nil, invalidLoginError
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, invalidUrlError
	}

	newAcc := AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}

	if password == "" {
		// shorthand can be used only if embedded struct was declared by short way
		newAcc.generatePassword(12) // is similar to
		// newAcc.Account.generatePassword(12)
		// newAcc.login = no_password_user
	}

	return &newAcc, nil
}
