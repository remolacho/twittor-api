package user

import "os/user"

type IUser interface {
	Create(u *user.User) (*user.User, error)
	ExistsByEmail(email string) bool
}
