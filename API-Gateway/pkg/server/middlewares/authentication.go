package middlewares

import (
	"log"
	"net/http"
	"strconv"
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
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error":"access token expired"})
	}

	User_id,User_role,err := jwttoken.GetRoleAndIDFromToken(tokenString)
	if err != nil{
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error":"token tampered"})
	}
	if User_role != "client" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error":"Unauthorised User"})
	}
	u_id,_:=strconv.Atoi(User_id)

	c.Locals("User_id",int64(u_id))
	c.Locals("User_role",User_role)

	log.Println("MW: User Authorized")
	return c.Next()
}

func AuthFreelancer(c *fiber.Ctx) error{
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var secretKey = viper.GetString("ATokenSecret")

	isValid, _ := jwttoken.IsValidAccessToken(secretKey, tokenString)
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error":"Not Authorised"})
	}

	User_id,User_role,err := jwttoken.GetRoleAndIDFromToken(tokenString)
	if err != nil{
		return c.Status(http.StatusUnauthorized).JSON("jwt token tampered")
	}
	if User_role != "freelancer" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error":"unauthorized user"})
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
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error":"Not Authorised"})
	}

	user_id,user_role,err := jwttoken.GetRoleAndIDFromToken(tokenString)
	if err != nil{
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error":"jwt token tampered"})
	}
	if user_role != "admin" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error":"unauthorized user"})
	}

	c.Set("user_id", user_id)
	c.Set("user_role",user_role)

	log.Println("MW: User Authorized")
	return c.Next()
}
