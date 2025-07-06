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
		option := promptData[string]([]string{
			"Меню:",
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Введите опцию",
		})

		switch option {
		case "1":
			createAccount(vault)
		case "2":
			findAccounts(vault)
		case "3":
			deleteAccount(vault)
		case "4":
			break Menu
		default:
			output.PrintError("Введена неизвестная опция")
		}
	}
}

func findAccounts(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите url"})

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
	login := promptData[string]([]string{"Введите логин"})
	url := promptData[string]([]string{"Введите url"})

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
	login := promptData[string]([]string{"Введите логин"})
	password := promptData[string]([]string{"Введите пароль"})
	urlString := promptData[string]([]string{"Введите url"})

	myAccount, err := account.NewAccount(login, password, urlString)
	if err != nil {
		output.PrintError(err)
		return
	}

	vault.AddAccount(myAccount)

	color.HiBlue("Аккаунт успешно добавлен")
}

func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}

	var res string
	_, err := fmt.Scanln(&res)
	if err != nil {
		return ""
	}
	return res
}
