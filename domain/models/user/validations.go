package user

import "golang.org/x/crypto/bcrypt"

func (u *User) EmailPresent() bool {
	return len(u.Email) > 0
}

func (u *User) PasswordSizeValid(size int) bool {
	return len(u.Password) >= size
}

func (u *User) EncryptPassword() (string, error) {
	cost := 8
	slicePassword := []byte(u.Password)
	bytes, err := bcrypt.GenerateFromPassword(slicePassword, cost)
	return string(bytes), err
}

func (u *User) DecodePassword(password string) error {
	passwordBytes := []byte(password)
	passwordDB := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	return err
}
