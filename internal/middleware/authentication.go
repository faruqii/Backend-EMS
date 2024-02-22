package middleware

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationMiddleware interface {
	Authenticate() error
}

type authenticationMiddleware struct {
	tokenRepository repositories.TokenRepository
}

func NewAuthenticationMiddleware(tokenRepository repositories.TokenRepository) *authenticationMiddleware {
	return &authenticationMiddleware{
		tokenRepository: tokenRepository,
	}
}

// Authenticate is a middleware that checks if the user is authenticated implements fiber's middleware interface
func (m *authenticationMiddleware) Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		user, err := m.tokenRepository.FindUserByToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		c.Locals("user", user)
		return c.Next()
	}
}

// Authorization is a middleware that checks if user's roles contain given target role on the param
func (m *authenticationMiddleware) Authorization(targetRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(string)
		if user != targetRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden",
			})
		}
		return c.Next()
	}
}