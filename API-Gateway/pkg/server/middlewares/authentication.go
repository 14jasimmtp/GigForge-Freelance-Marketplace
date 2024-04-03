package middlewares

import (
	"log"
	"net/http"
	"strings"

	jwttoken "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/utils/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func AuthClient(c *fiber.Ctx) error {
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var secretKey = viper.GetString("ATokenSecret")

	isValid, _ := jwttoken.IsValidAccessToken(secretKey, tokenString)
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON("Not Authorised")
	}

	User_id,User_role,err := jwttoken.GetRoleAndIDFromToken(tokenString)
	if err != nil{
		return c.Status(http.StatusUnauthorized).JSON("jwt token tampered")
	}
	if User_role != "client" {
		return c.Status(http.StatusForbidden).JSON("In wrong way")
	}

	c.Locals("User_id",User_id)
	c.Locals("User_role",User_role)

	log.Println("MW: User Authorized")
	return c.Next()
}

func AuthFreelancer(c *fiber.Ctx) error{
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var secretKey = viper.GetString("ATokenSecret")

	isValid, _ := jwttoken.IsValidAccessToken(secretKey, tokenString)
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON("Not Authorised")
	}

	User_id,User_role,err := jwttoken.GetRoleAndIDFromToken(tokenString)
	if err != nil{
		return c.Status(http.StatusUnauthorized).JSON("jwt token tampered")
	}
	if User_role != "freelancer" {
		return c.Status(http.StatusForbidden).JSON("In wrong way")
	}

	c.Locals("User_id", User_id)
	c.Locals("User_role",User_role)

	log.Println("MW: User Authorized")
	return c.Next()
}

func AuthAdmin(c *fiber.Ctx) error{
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var secretKey = viper.GetString("ATokenSecret")

	isValid, _ := jwttoken.IsValidAccessToken(secretKey, tokenString)
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON("Not Authorised")
	}

	user_id,user_role,err := jwttoken.GetRoleAndIDFromToken(tokenString)
	if err != nil{
		return c.Status(http.StatusUnauthorized).JSON("jwt token tampered")
	}
	if user_role != "admin" {
		return c.Status(http.StatusForbidden).JSON("In wrong way")
	}

	c.Set("user_id", user_id+"1")
	c.Set("user_role",user_role)

	log.Println("MW: User Authorized")
	return c.Next()
}
