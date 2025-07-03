package account

import (
	"encoding/json"
	"strings"
	"time"
	
	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func NewVault(db Db) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Failed to read data.json")
	}
	
	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, jsonError
	}

	return file, nil
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать в JSON")
	}

	vault.db.Write(data)
}

func (vault *VaultWithDb) AddAccount(account *Account) {
	vault.Accounts = append(vault.Accounts, *account)
	vault.save()
}

func (vault *VaultWithDb) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			accounts = append(accounts, acc)
		}
	}

	return accounts
}

func (vault *VaultWithDb) FindAccount(login string, url string) *Account {
	for _, account := range vault.Accounts {
		if account.Login == login && account.Url == url {
			return &account
		}
	}

	return nil
}

func (vault *VaultWithDb) DeleteAccount(login string, url string) {
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
