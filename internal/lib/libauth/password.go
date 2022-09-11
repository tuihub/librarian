package libauth

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func (a *Auth) GeneratePassword(password string) (string, error) {
	res, err := scrypt.Key(
		[]byte(password),
		[]byte(a.config.Salt),
		1<<14, 8, 1, 32) //nolint:gomnd // default crypt settings
	if err == nil {
		return fmt.Sprintf("%x", res), nil
	}
	return "", err
}
