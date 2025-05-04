package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var availableLetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*@!_.")

var (
	InvalidUrlError   = errors.New("url is invalid")
	InvalidLoginError = errors.New("login is empty")
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct { // composition struct
	createdAt time.Time
	updatedAt time.Time
	account // simple and preferred way
	// acc  account --> second and longer way to use composition
}

func (acc *account) outputPassword() { // (acc account) is a copy of struct instance
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) { // acc *account is a pointer to instance
	res := make([]rune, n)

	for i := range res {
		res[i] = availableLetterRunes[rand.IntN(len(availableLetterRunes)-1)]
	}

	acc.password = string(res)
}

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, InvalidLoginError
	}
	
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, InvalidUrlError
	}
	
	newAcc := accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	
	if password == "" {
		// shorthand can be used only if embedded struct was declared by short way
		newAcc.generatePassword(12) // is similar to
		// newAcc.account.generatePassword(12)
		// newAcc.login = no_password_user
	}
	
	return &newAcc, nil
}

func newAccount(login, password, urlString string) (*account, error) { // account's constructor
	if login == "" {
		return nil, InvalidLoginError
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, InvalidUrlError
	}

	newAcc := account{
		login:    login,
		password: password,
		url:      urlString,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return &newAcc, nil
}

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	urlString := promptData("Введите url: ")

	myAccount, err := newAccountWithTimeStamp(login, password, urlString)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAccount.outputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	_, err := fmt.Scanln(&res)
	if err != nil {
		return ""
	}
	return res
}
