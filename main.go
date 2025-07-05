package main

import (
	"fmt"

	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("__Менеджер паролей__")

	vault := account.NewVault(files.NewJsonDb("data.json"))

Menu:
	for {
		option := getMenu()

		switch option {
		case 1:
			createAccount(vault)
		case 2:
			findAccounts(vault)
		case 3:
			deleteAccount(vault)
		case 4:
			break Menu
		default:
			output.PrintError("Введена неизвестная опция")
		}
	}
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

func findAccounts(vault *account.VaultWithDb) {
	url := promptData("Введите url: ")

	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
		return
	}

	color.Magenta("Результаты поиска:\n\n")
	for i, acc := range accounts {
		color.Cyan(`-- Аккаунт %d --`, i+1)
		acc.Output()
		fmt.Println()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин: ")
	url := promptData("Введите url: ")

	acc := vault.FindAccount(login, url)
	if acc == nil {
		output.PrintError("Аккаунт не найден")
		output.PrintError("Не удалось удалить аккаунт")
		return
	}

	vault.DeleteAccount(login, url)

	color.Green("Аккаунт успешно удален")
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	urlString := promptData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, urlString)
	if err != nil {
		output.PrintError(err)
		return
	}

	vault.AddAccount(myAccount)

	color.HiBlue("Аккаунт успешно добавлен")
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
