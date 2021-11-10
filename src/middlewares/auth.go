package middlewares

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	jwt.StandardClaims
	Scope string
	Email string
}

func IsAuth(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	payload := token.Claims.(*CustomClaims)

	if err != nil || !token.Valid {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	IsAmbasador := strings.Contains(c.Path(), "/api/ambasador")

	if payload.Scope == "admin" && IsAmbasador || payload.Scope == "ambasador" && !IsAmbasador {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Unauthorised",
		})
	}

	return c.Next()
}

func GetUserId(c *fiber.Ctx) (uint, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	payload := token.Claims.(*CustomClaims)

	userId, _ := strconv.Atoi(payload.Subject)
	return uint(userId), nil
}

func GenerateToken(userId uint, scope string) (string, error) {
	var payload = CustomClaims{
		Email: "stajesad@gmail.com",
		Scope: scope,
	}

	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	payload.Subject = strconv.Itoa(int(userId))

	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
}
