package main

import (
	"fmt"
	
	"demo/password/account"
	"demo/password/files"
)

const storageFileName = "data.json"

func main() {
	fmt.Println("__Менеджер паролей__")

Menu:
	for {
		option := getMenu()

		switch option {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		case 4:
			break Menu
		default:
			fmt.Println("Введена неверная опция")
		}
	}

	fmt.Println("Успешный выход из приложения")
}

func getMenu() int {
	fmt.Println(
		"Меню:\n" +
			"1. Создать аккаунт\n" +
			"2. Найти аккаунт\n" +
			"3. Удалить аккаунт\n" +
			"4. Выход",
	)

	fmt.Print("Введите опцию: ")
	var option int
	_, err := fmt.Scanln(&option)
	if err != nil {
		return -1
	}

	return option
}

func findAccount() {
	fmt.Println("Account data...")
}

func deleteAccount() {
	fmt.Println("Account deleted")
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

	existingData := files.ReadFile(storageFileName)
	vault, err := account.GetVault(existingData)
	if err != nil {
		vault = account.NewVault()
	}
	vault.AddAccount(myAccount)

	file, e := vault.ToBytes()
	if e != nil {
		fmt.Println(e)
		return
	}

	files.WriteFile(file, storageFileName)
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
