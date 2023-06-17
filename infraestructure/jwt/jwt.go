package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
	"time"
	"twittor-api/domain/models/user"
)

type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id, omitempty" json:"_id,omitempty"`
	jwt.StandardClaims
}

type Jwt struct {
	Secret []byte
}

func NewJwt() *Jwt {
	return &Jwt{
		[]byte(os.Getenv("JWT_PASSWORD")),
	}
}

func (j *Jwt) Encode(u user.User) (string, error) {
	if u.Email == "" {
		return "", errors.New("the user is null")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload(u))
	return token.SignedString(j.Secret)
}

func (j *Jwt) Decode(jwtKey string) (*Claim, error) {
	claim := &Claim{}
	splitToken := strings.Split(jwtKey, " ")

	if len(splitToken) != 2 {
		return nil, errors.New("the jwt is not valid")
	}

	token := strings.TrimSpace(splitToken[1])
	result, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})

	if err != nil {
		return nil, errors.New("the jwt is not valid " + err.Error())
	}

	if !result.Valid {
		return nil, errors.New("the jwt is not valid")
	}

	return claim, nil
}

func payload(u user.User) jwt.MapClaims {
	return jwt.MapClaims{
		"email":     u.Email,
		"name":      u.Name,
		"lastname":  u.LastName,
		"biography": u.Biography,
		"sideWeb":   u.SideWeb,
		"location":  u.Location,
		"_id":       u.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
}
