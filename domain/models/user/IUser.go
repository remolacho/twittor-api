package user

type IUser interface {
	Create(u *User) (*User, error)
	ExistsByEmail(email string) bool
	FindByEmail(email string) (*User, error)
	Find(ID string) (*User, error)
	Update(u *User) (*User, error)
	AllUnFollowed(ID string, page int64, searchTerm string) ([]User, error)
	AllFollowed(ID string, page int64, searchTerm string) ([]User, error)
}
