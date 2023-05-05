package user_repository

type UserRepository struct{}

func New() *UserRepository {
	return &UserRepository{}
}
