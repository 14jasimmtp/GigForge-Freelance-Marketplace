package jwt

import (
	"time"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/admin-svc/pkg/domain"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type ClientClaims struct {
	jwt.RegisteredClaims
	User_id int
	Email   string
	Role    string
}

func AdminTokenGenerate(user *domain.Admin) (string, error) {
	claims := ClientClaims{
		User_id: int(user.ID),
		Email:   user.Email,
		Role:    "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "GigForge",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString([]byte(viper.GetString("ATokenSecret")))
	if err != nil {
		return "", err
	}
	return TokenString, nil
}
