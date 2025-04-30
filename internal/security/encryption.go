package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

var ENCRYPTION_KEY []byte

func InitEncryptionKey() {
	keySecrect := os.Getenv("ENCRYPTION_KEY")
	ENCRYPTION_KEY = []byte(keySecrect)
}

func Encrypt(key string) (string, error) {

	block, err := aes.NewCipher(ENCRYPTION_KEY)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(key), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(secret string) (string, error) {
	
	cipherTxt, _ := base64.StdEncoding.DecodeString(secret)

	block, _ := aes.NewCipher(ENCRYPTION_KEY)
    aesGCM, _ := cipher.NewGCM(block)

    nonceSize := aesGCM.NonceSize()
    nonce, cipherData := cipherTxt[:nonceSize], cipherTxt[nonceSize:]

    plainText, err := aesGCM.Open(nil, nonce, cipherData, nil)
    if err != nil {
        return "", err
    }

    return string(plainText), nil
}