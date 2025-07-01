package main

import (
	"fmt"
	
	"demo/password/account"
)

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
	login := promptData("Введите логин: ")
	url := promptData("Введите url: ")
	
	vault := account.NewVault()

	acc := vault.FindAccount(login, url)
	if acc == nil {
		fmt.Println("Аккаунт не найден")
		return
	}

	fmt.Printf(
		`Аккаунт:
логин: %s
пароль: %s
url: %s
`,
		acc.Login, acc.Password, acc.Url,
	)
}

func deleteAccount() {
	login := promptData("Введите логин: ")
	url := promptData("Введите url: ")
	
	vault := account.NewVault()
	
	acc := vault.FindAccount(login, url)
	if acc == nil {
		fmt.Println("Аккаунт не найден")
		return
	}
	
	vault.DeleteAccount(login, url)
	
	fmt.Println("Аккаунт успешно удален")
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

	vault := account.NewVault()
	vault.AddAccount(myAccount)
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
