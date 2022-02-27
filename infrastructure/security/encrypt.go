package security

import "golang.org/x/crypto/bcrypt"

type EncryptVal string

func Encrypt(val string) (EncryptVal, error) {
	newVal, err := bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)
	return EncryptVal(newVal), err
}

func (ev EncryptVal) Verify(val string) error {
	return bcrypt.CompareHashAndPassword([]byte(ev), []byte(val))
}
