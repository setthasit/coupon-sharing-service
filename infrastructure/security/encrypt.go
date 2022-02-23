package security

import "golang.org/x/crypto/bcrypt"

type EncryptVal []byte

func Encrypt(val string) (EncryptVal, error) {
	return bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)
}

func (ev EncryptVal) Verify(val string) error {
	return bcrypt.CompareHashAndPassword(ev, []byte(val))
}
