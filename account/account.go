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
	jsonError = errors.New("failed to convert data to JSON")
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.Password)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = availableLetterRunes[rand.IntN(len(availableLetterRunes)-1)]
	}

	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (
	*Account, error,
) {
	if login == "" {
		return nil, invalidLoginError
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, invalidUrlError
	}

	newAcc := Account{
		Login:    login,
		Password: password,
		Url:      urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return &newAcc, nil
}
