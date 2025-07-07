package main

import (
	"fmt"
	"strings"

	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountsByUrl,
	"3": findAccountsByLogin,
	"4": deleteAccount,
}

var menuLabels = []string{
	"Меню:",
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Введите опцию",
}

func main() {
	fmt.Println("__Менеджер паролей__")

	vault := account.NewVault(files.NewJsonDb("data.json"))

	for {
		option := promptData(menuLabels...)

		if option == "5" {
			break
		}

		if menu[option] == nil {
			output.PrintError("Введена неизвестная опция")
			continue
		}

		menu[option](vault)
	}
}

func findAccountsByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")

	accounts := vault.FindAccounts(
		url, func(acc account.Account, str string) bool {
			return strings.Contains(strings.ToLower(acc.Url), strings.ToLower(str))
		},
	)

	showFindResults(accounts)
}

func findAccountsByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска")

	accounts := vault.FindAccounts(
		login, func(acc account.Account, str string) bool {
			return strings.Contains(strings.ToLower(acc.Login), strings.ToLower(str))
		},
	)

	showFindResults(accounts)
}

func showFindResults(accounts []account.Account) {
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
	login := promptData("Введите логин")
	url := promptData("Введите url")

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
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	urlString := promptData("Введите url")

	myAccount, err := account.NewAccount(login, password, urlString)
	if err != nil {
		output.PrintError(err)
		return
	}

	vault.AddAccount(myAccount)

	color.HiBlue("Аккаунт успешно добавлен")
}

func promptData(prompt ...string) string {
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
