package main

import (
	"fmt"
	"math/rand/v2"
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

func newAccount(login, url string) *account {
	acc := account{
		login: login,
		url: url,
	}
	acc.generatePassword(12)
	
	return &acc
}

var availableLetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*@!_.")

func main() {
	login := promptData("Введите логин: ")
	url := promptData("Введите url: ")

	myAccount := newAccount(login, url)
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
