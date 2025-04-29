package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите url: ")
	
	myAccount := account{
		login,    // ordering is important
		password, // passing the parameters in same order as in struct
		url,
	}
	
	outputPassword(&myAccount)
}

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)  // (*acc).login is similar as acc.login
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
