package hash

import "golang.org/x/crypto/bcrypt"

func IsEqualWithHash(original string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(original))
	return err == nil
}

func ToHashString(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(bytes), err
}
