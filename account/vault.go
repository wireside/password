package account

import (
	"encoding/json"
	"strings"
	"time"
	
	"demo/password/encrypter"
	"demo/password/output"
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
	enc encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	encryptedFile, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
			enc: enc,
		}
	}

	var vault Vault
	data := enc.Decrypt(encryptedFile)
	err = json.Unmarshal(data, &vault)
	color.Cyan("Найдено аккаунтов: %d", len(vault.Accounts))
	if err != nil {
		output.PrintError("Не удалось прочитать JSON")
	}

	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc: enc,
	}
}

func (vault *VaultWithDb) AddAccount(account *Account) {
	vault.Accounts = append(vault.Accounts, *account)
	vault.save()
}

func (vault *VaultWithDb) FindAccountsByUrl(url string) []Account {
	return vault.FindAccounts(
		url, func(acc Account, str string) bool {
			return strings.Contains(strings.ToLower(acc.Url), strings.ToLower(str))
		},
	)
}

func (vault *VaultWithDb) FindAccountsByLogin(login string) []Account {
	return vault.FindAccounts(
		login, func(acc Account, str string) bool {
			return strings.Contains(strings.ToLower(acc.Login), strings.ToLower(str))
		},
	)
}

func (vault *VaultWithDb) FindAccounts(
	param string,
	checker func(Account, string) bool,
) []Account {
	var accounts []Account
	for _, acc := range vault.Accounts {
		isMatched := checker(acc, param)
		if isMatched {
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
	encryptedData := vault.enc.Encrypt(data)
	if err != nil {
		output.PrintError("Не удалось преобразовать в JSON")
	}
	
	vault.db.Write(encryptedData)
}
