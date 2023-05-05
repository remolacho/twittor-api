package user

type IUser interface {
	Create(u *User) (*User, error)
	ExistsByEmail(email string) bool
	FindByEmail(email string) (*User, error)
	Find(ID string) (*User, error)
}
