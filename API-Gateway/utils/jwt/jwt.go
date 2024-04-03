package jwttoken

import (
	"fmt"
	"strconv"
	"time"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/auth"
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

func IsValidVerifyToken(secretKey, tokenString string) (bool, *TempTokenClaims) {

	// Parse jwt token with custom claims
	token, err := jwt.ParseWithClaims(tokenString, &TempTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	// Check if token is valid
	if err != nil || !token.Valid {
		fmt.Println("Error occurred while parsing token:", err)
		return false, nil
	}

	// Assign parsed data from token to claims
	if claims, ok := token.Claims.(*TempTokenClaims); ok && token.Valid {

		// Check if token is expired
		if claims.ExpiresAt.Before(time.Now()) {
			fmt.Println("token expired")
			return false, nil
		}

		return true, claims

	} else {
		fmt.Println("Error occurred while decoding token:", err)
		return false, nil
	}
}

func IsValidAccessToken(secretKey, tokenString string) (bool, *AccessTokenClaims) {

	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Error occurred while parsing token:", err)
		return false, nil
	}

	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {

		// Check if token is expired
		if claims.ExpiresAt.Before(time.Now()) {
			fmt.Println("token expired")
			return false, nil
		}

		return true, claims

	} else {
		fmt.Println("Error occurred while decoding token:", err)
		return false, nil
	}
}

func GetRoleAndIDFromToken(Token string) (string, string, error) {

	fmt.Println(viper.GetString("ATokenSecret"))
	TokenUnpacked, err := jwt.ParseWithClaims(Token, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("1")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("ATokenSecret")), nil
	})
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	if claims, ok := TokenUnpacked.Claims.(*AccessTokenClaims); ok && TokenUnpacked.Valid {
		fmt.Println(claims.User_id,claims.Role,claims.Email)
		return strconv.Itoa(claims.User_id), claims.Role, nil
	}

	fmt.Println("3")
	return "", "", fmt.Errorf("invalid token")
}
