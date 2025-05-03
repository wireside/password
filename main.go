package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

var availableLetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*@!_.")

var InvalidUrlError = errors.New("URL is invalid, please enter the correct URL")

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

func newAccount(login, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, InvalidUrlError
	}

	acc := account{
		login: login,
		url:   urlString,
	}
	acc.generatePassword(12)

	return &acc, nil
}

func main() {
	login := promptData("Введите логин: ")
	urlString := promptData("Введите urlString: ")

	myAccount, err := newAccount(login, urlString)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	_, err := fmt.Scan(&res)
	if err != nil {
		return ""
	}
	return res
}
