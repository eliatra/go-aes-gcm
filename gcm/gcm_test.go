package gcm

import (
	"bytes"
	"testing"
)

var key1 = []byte("12345678123456781234567812345678")
var key2 = []byte("11111111111111111111111111111111")
var plainTextBytes = []byte("this is plaintext")
var plainTextString = "this is plaintext"
var aad = []byte("some aad")

func TestAesGcm(t *testing.T) {

	a, err := Encrypt(plainTextBytes, key1, aad)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	b, err := Decrypt(a, key1, aad)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !bytes.Equal(b, plainTextBytes) {
		t.Fatalf("Expected %v == %v", b, plainTextBytes)
	}
}

func TestAesGcmWrongKey(t *testing.T) {

	a, err := Encrypt(plainTextBytes, key1, aad)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	_, err = Decrypt(a, key2, aad)
	if err == nil {
		t.Fatalf("Unexpected error %v", err)
	}
}

func TestAesGcmWrongAad(t *testing.T) {

	a, err := Encrypt(plainTextBytes, key1, aad)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	_, err = Decrypt(a, key1, nil)
	if err == nil {
		t.Fatalf("Unexpected error %v", err)
	}

}

func TestAesGcmString(t *testing.T) {

	a, err := EncryptStringToString(plainTextString, key1, aad)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	b, err := DecryptFromStringToString(a, key1, aad)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if b != plainTextString {
		t.Fatalf("Expected %v == %v", b, plainTextBytes)
	}
}

func TestAesGcmStringWrongAad(t *testing.T) {

	a, err := EncryptStringToString(plainTextString, key1, aad)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	_, err = DecryptFromStringToString(a, key1, key2)
	if err == nil {
		t.Fatalf("Unexpected error %v", err)
	}

}
