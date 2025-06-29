package main

import (
	"fmt"

	"demo/password/account"
	"demo/password/files"
)

func main() {
	createAccount()
}

func createAccount() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	urlString := promptData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, urlString)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println(err)
		return
	}

	files.WriteFile(file, "data.json")
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
