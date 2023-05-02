package jwt_service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
	"twittor-api/domain/models/user"
)

type JwtService struct {
	*user.User
	Secret []byte
}

func NewJwt(u *user.User) *JwtService {
	return &JwtService{
		u,
		[]byte(os.Getenv("JWT_PASSWORD")),
	}
}

func (j *JwtService) Encode() (string, error) {
	if j.User == nil {
		return "", errors.New("the user is null")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload(j.User))
	return token.SignedString(j.Secret)
}

func payload(u *user.User) jwt.MapClaims {
	return jwt.MapClaims{
		"email":     u.Email,
		"name":      u.Name,
		"lastname":  u.LastName,
		"biography": u.Biography,
		"site":      u.SiteWeb,
		"location":  u.Location,
		"_id":       u.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
}
