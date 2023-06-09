package user_service

import (
	"twittor-api/domain/models/user"
	jwtService "twittor-api/infraestructure/jwt"
)

type ResponseJwt struct {
	Token string `json:"token,omitempty"`
}

type UserLoginService struct {
	RepositoryUser user.IUser
}

func NewLogin(repository user.IUser) *UserLoginService {
	return &UserLoginService{
		repository,
	}
}

func (s *UserLoginService) Login(email, password string) (ResponseJwt, bool) {
	var response ResponseJwt

	u, err := s.RepositoryUser.FindByEmail(email)

	if err != nil {
		return response, false
	}

	err = u.DecodePassword(password)

	if err != nil {
		return response, false
	}

	serviceJwt := jwtService.NewJwt()
	response.Token, err = serviceJwt.Encode(*u)

	if err != nil {
		return response, false
	}

	response.Token = "Bearer " + response.Token

	return response, true
}
