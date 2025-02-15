package middlewares

import (
	"strings"

	"github.com/devs-group/driplet/api/auth"
	"github.com/devs-group/driplet/api/repositories"
	"github.com/gofiber/fiber/v2"
)

type AuthConfig struct {
	TokenValidator  *auth.TokenValidator
	UsersRepository *repositories.UsersRepository
}

func RequireAuth(config AuthConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Authorization header required",
			})
		}

		// Extract token from "Bearer <token>"
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid authorization format",
			})
		}

		// Validate Google token
		claims, err := config.TokenValidator.ValidateGoogleToken(token)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Get or create user
		user, err := config.UsersRepository.FindByEmail(claims.Email)
		if err != nil {
			// If user doesn't exist, create them
			user = &repositories.User{
				Email: claims.Email,
			}
			if err := config.UsersRepository.Create(user); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "Failed to create user",
				})
			}
		}

		// Attach user to context
		c.Locals("user", user)

		return c.Next()
	}
}
