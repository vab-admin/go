package hash

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// PasswordHash
// @param password
// @date 2022-09-10 17:38:06
func PasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return hex.EncodeToString(bytes)
}

// PasswordVerify
// @param hash
// @param password
// @date 2022-09-10 17:38:04
func PasswordVerify(hash, password string) bool {
	h, err := hex.DecodeString(hash)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(h, []byte(password))
	return err == nil
}
