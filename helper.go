package ftech

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func DoHash(text string) string {
	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted)
}

func DoHashUsingSalt(text string) string {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted)
}
