package user

import "os/user"

type IUser interface {
	Create(u *user.User) (error, *user.User)
	ExistsByEmail(email string) bool
}
