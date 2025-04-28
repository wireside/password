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
	
	account1 := account{
		login,       // ordering is important
		password, // passing the parameters in same order as in struct
		url,
	}
	
	account2 := account{
		password: password, // property name is important
		login:    login,    // you can pass the parameters in any order you want
		url:      url,
	}
	
	account3 := account{
		password: password, // without login --> login sets to empty string ""
		url:      url,
	}
	
	emptyAccount := account{} // empty instance which includes default null values
	
	fmt.Println(account1, account2, account3, emptyAccount)
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
