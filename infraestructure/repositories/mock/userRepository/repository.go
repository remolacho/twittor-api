package userRepository

type UserRepository struct{}

func New() *UserRepository {
	return &UserRepository{}
}
