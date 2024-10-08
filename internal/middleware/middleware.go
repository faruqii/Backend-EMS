package middleware

import (
	"strings"

	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationMiddleware interface {
	Authenticate() fiber.Handler
}

type AuthorizationMiddleware interface {
	Authorization(targetRole string) fiber.Handler
}

type Middleware struct {
	tokenRepository repositories.TokenRepository
	roleRepository  repositories.RoleRepository
}

func NewMiddleware(tokenRepository repositories.TokenRepository, roleRepository repositories.RoleRepository) *Middleware {
	return &Middleware{
		tokenRepository: tokenRepository,
		roleRepository:  roleRepository,
	}
}

type MiddlewareError struct {
	Message    string
	StatusCode int
}

func TestMode() fiber.Handler {
	return func(c *fiber.Ctx) error {
		testMode := c.Query("test") == "true" || c.Get("X-Test-Mode") == "true"
		c.Locals("testMode", testMode)
		return c.Next()
	}
}

// Authenticate middleware checks if the user is authenticated and sets user data in context locals
func (m *Middleware) Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}

		// Extract bearer token from Authorization header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}

		// Fetch user from token
		user, err := m.tokenRepository.FindUserByToken(token)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}

		// Fetch user's role name from the repository based on user's ID
		userRoleName, err := m.tokenRepository.FindRoleTypeBasedOnToken(token)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		c.Locals("user", user)
		c.Locals("role", userRoleName) // Set user's role name in locals

		// Set bearer token as a custom header
		c.Set("auth", token)

		return c.Next()
	}
}

// Authorization middleware checks if the user has any of the required roles
func (m *Middleware) Authorization(targetRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, ok := c.Locals("role").(string)
		if !ok {
			return fiber.NewError(fiber.StatusForbidden, "Forbidden")
		}

		// Check if the user has any of the required roles
		for _, role := range targetRoles {
			if userRole == role {
				return c.Next()
			}
		}

		return fiber.NewError(fiber.StatusForbidden, "Forbidden")
	}
}
