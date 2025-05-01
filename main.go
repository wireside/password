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

var availableLetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*@!_.")

func main() {
	login := promptData("Введите логин: ")
	password := generatePassword(8)
	url := promptData("Введите url: ")

	myAccount := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(&myAccount)
}

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url) // (*acc).login is similar as acc.login
	// fmt.Println((*acc).login, (*acc).password, (*acc).url) without shorthand
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

func generatePassword(n int) string {
	res := make([]rune, n)

	for i := range res {
		index := rand.IntN(len(availableLetterRunes) - 1)
		letter := availableLetterRunes[index]
		res[i] = letter
	}

	return string(res)
}
