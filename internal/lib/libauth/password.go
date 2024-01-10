package libauth

import (
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
)

func (a *Auth) GeneratePassword(password string) (string, error) {
	res, err := scrypt.Key(
		[]byte(password),
		[]byte(a.config.GetSalt()),
		1<<14, 8, 5, 32) //nolint:gomnd // based on https://tobtu.com/minimum-password-settings/
	if err == nil {
		return hex.EncodeToString(res), nil
	}
	return "", err
}
