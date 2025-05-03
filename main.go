package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

var availableLetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*@!_.")

var (
	InvalidUrlError = errors.New("url is invalid")
	InvalidLoginError = errors.New("login is empty")
)

type account struct {
	login    string
	password string
	url      string
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

func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, InvalidLoginError
	}
	
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, InvalidUrlError
	}

	newAcc := account{
		login: login,
		password: password,
		url:   urlString,
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

	myAccount, err := newAccount(login, password, urlString)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAccount.outputPassword()
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
