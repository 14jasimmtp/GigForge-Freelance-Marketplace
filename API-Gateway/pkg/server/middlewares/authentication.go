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

	claims, err := jwttoken.IsValidAccessToken(secretKey, tokenString)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	if claims.Role != "client" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Unauthorised User"})
	}

	c.Locals("User_id", claims.User_id)
	c.Locals("User_role", claims.Role)
	log.Println("MW: User Authorized")
	return c.Next()
}

func AuthFreelancer(c *fiber.Ctx) error {
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var secretKey = viper.GetString("ATokenSecret")

	Claims, err := jwttoken.IsValidAccessToken(secretKey, tokenString)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Not Authorised"})
	}

	if Claims.Role != "freelancer" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "unauthorized user"})
	}

	c.Locals("User_id", Claims.User_id)
	c.Locals("User_role", Claims.Role)

	log.Println("MW: User Authorized")
	return c.Next()
}

func AuthChat(c *fiber.Ctx) error {
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var secretKey = viper.GetString("ATokenSecret")

	Claims, err := jwttoken.IsValidAccessToken(secretKey, tokenString)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Not Authorised"})
	}

	c.Locals("User_id", Claims.User_id)
	c.Locals("User_role", Claims.Role)

	log.Println("MW: User Authorized")
	return c.Next()
}

func AuthAdmin(c *fiber.Ctx) error {
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var secretKey = viper.GetString("ATokenSecret")

	Claims, err := jwttoken.IsValidAccessToken(secretKey, tokenString)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Not Authorised"})
	}

	c.Locals("user_id", Claims.User_id)
	c.Locals("user_role", Claims.Role)

	log.Println("MW: User Authorized")
	return c.Next()
}
