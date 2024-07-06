package gcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

var salt = []byte("gah78j1rTTnma675")

func Decrypt(cipherText, secretKey, aad []byte) ([]byte, error) {
	aes, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce := cipherText[:nonceSize]
	ciphertextWithoutNonce := cipherText[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, ciphertextWithoutNonce, aad)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func Encrypt(plainText, secretKey, aad []byte) ([]byte, error) {
	aes, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	r, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	if r != gcm.NonceSize() {
		return nil, errors.New("nonce length invalid")
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, aad)

	return cipherText, nil
}

func EncryptFile(plainTextFile, cipherTextFile string, secretKey []byte) error {
	plainText, err := os.ReadFile(plainTextFile)
	if err != nil {
		return err
	}

	cipherText, err := Encrypt(plainText, secretKey, nil)

	if err != nil {
		return err
	}

	if err := os.WriteFile(cipherTextFile, cipherText, 0600); err != nil {
		return err
	}

	return nil
}

func DecryptFile(cipherTextFile, plainTextFile string, secretKey []byte) error {
	cipherText, err := os.ReadFile(cipherTextFile)
	if err != nil {
		return err
	}

	plainText, err := Decrypt(cipherText, secretKey, nil)

	if err != nil {
		return err
	}

	if err := os.WriteFile(plainTextFile, plainText, 0600); err != nil {
		return err
	}

	return nil
}

func CreateKeyFromPassword(password string) ([]byte, error) {
	if len(password) < 8 {
		return nil, errors.New("password too short")
	}

	return pbkdf2.Key([]byte(password), salt, 210000, 32, sha512.New), nil
}
