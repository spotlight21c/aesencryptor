package aesencryptor_test

import (
	"encoding/base64"
	"testing"

	"github.com/spotlight21c/aesencryptor"
)

func TestEncrypt(t *testing.T) {
	tables := map[string]string{
		"This-is-plain-text": "ACFbo72am8SBFGTHgiUlDygR1jYdLLrCWPeFJ2BlyRU=",
	}

	key := "1234567890abcdef"

	for plain, result := range tables {
		encValue, _ := aesencryptor.Encrypt(plain, key)
		encodedString := base64.StdEncoding.EncodeToString(encValue)
		if encodedString != result {
			t.Errorf("%s expected %s, but %s", plain, result, encodedString)
		}
	}
}
