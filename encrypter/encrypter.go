package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
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

func (enc *Encrypter) Encrypt(plaintText []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	
	nonce := make([]byte, aesGCM.NonceSize())
	
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	
	return aesGCM.Seal(nonce, nonce, plaintText, nil)
}

func (enc *Encrypter) Decrypt(encryptedText []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	
	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := encryptedText[:nonceSize], encryptedText[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	
	return plainText
}