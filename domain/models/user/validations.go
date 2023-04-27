package user

func (u *User) EmailPresent() bool {
	return len(u.Email) > 0
}

func (u *User) PasswordSizeValid(size int) bool {
	return len(u.Password) >= size
}
