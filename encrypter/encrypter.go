package encrypter

import (
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		panic("Не определена переменная окружения SECRET_KEY")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plaintStr string) string {
	return ""
}

func (enc *Encrypter) Decrypt(encryptedStr string) string {
	return ""
}