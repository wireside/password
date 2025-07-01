package account

import (
	"encoding/json"
	"strings"
	"time"

	"demo/password/files"
	"github.com/fatih/color"
)

const StorageFileName = "data.json"

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, jsonError
	}

	return file, nil
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать в JSON")
	}

	files.WriteFile(data, StorageFileName)
}

func NewVault() *Vault {
	file, err := files.ReadFile(StorageFileName)
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Failed to read data.json")
	}

	return &vault
}

func (vault *Vault) AddAccount(account *Account) {
	vault.Accounts = append(vault.Accounts, *account)
	vault.save()
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			accounts = append(accounts, acc)
		}
	}

	return accounts
}

func (vault *Vault) FindAccount(login string, url string) *Account {
	for _, account := range vault.Accounts {
		if account.Login == login && account.Url == url {
			return &account
		}
	}

	return nil
}

func (vault *Vault) DeleteAccount(login string, url string) {
	accounts := make([]Account, 0, len(vault.Accounts))

	for _, acc := range vault.Accounts {
		if acc.Login == login && acc.Url == url {
			continue
		}

		accounts = append(accounts, acc)
	}

	vault.Accounts = accounts
	vault.save()
}
