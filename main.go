package main

import (
	"fmt"
	
	"demo/password/account"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("__Менеджер паролей__")
	
	vault := account.NewVault()

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

func findAccounts(vault *account.Vault) {
	url := promptData("Введите url: ")

	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
		return
	}
	
	color.Magenta("Результаты поиска:\n")
	for i, acc := range accounts {
		color.Cyan(`-- Аккаунт %d --`, i + 1)
		acc.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	login := promptData("Введите логин: ")
	url := promptData("Введите url: ")
	
	acc := vault.FindAccount(login, url)
	if acc == nil {
		color.Red("Аккаунт не найден")
		color.Red("Не удалось удалить аккаунт")
		return
	}
	
	vault.DeleteAccount(login, url)
	
	color.Green("Аккаунт успешно удален")
}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	urlString := promptData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, urlString)
	if err != nil {
		fmt.Println(err)
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
