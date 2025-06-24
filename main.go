package main

import (
	"fmt"

	"demo/password/account"
	"demo/password/files"
)

func main() {
	files.WriteFile("Привет! Я файл", "file.txt")
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	urlString := promptData("Введите url: ")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, urlString)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAccount.OutputPassword()
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
