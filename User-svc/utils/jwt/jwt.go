package jwtoken

import (
	"fmt"
	"strings"
	"time"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/domain"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type TempTokenClaims struct {
	User *auth.UserSignupReq
	Role string
	jwt.RegisteredClaims
}

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	User_id int
	Email   string
	Role    string
}

func GenerateAccessToken(secret string, user *domain.UserModel) (string, error) {
	fmt.Println(user.ID,"user id")
	claims := AccessTokenClaims{
		User_id: int(user.ID),
		Email:   user.Email,
		Role:    user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "GigForge",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	AccessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return AccessToken, nil
}

// func GenerateRefreshToken(secret, email string) (string, error) {
// 	claims := jwt.MapClaims{
// 		"email": email,
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte(secret))
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

func GenerateTemporaryTokenToVerify(secret string, user *auth.UserSignupReq) (string, error) {
	claims := &TempTokenClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString([]byte(secret))
	fmt.Println(secret)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func FetchUserVerifyDetailsFromToken(token string) (*auth.UserSignupReq, error) {
	tokenString := strings.TrimPrefix(token, "Bearer ")
	TokenUnpacked, err := jwt.ParseWithClaims(tokenString, &TempTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("ATokenSecret")), nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if claims, ok := TokenUnpacked.Claims.(*TempTokenClaims); ok && TokenUnpacked.Valid {
		return claims.User, nil
	}

	return nil, nil
}
