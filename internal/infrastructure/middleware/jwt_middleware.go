package middleware

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mBuergi86/devseconnect/pkg/security"
	"github.com/rs/zerolog"
)

func JWTMiddleware(logger zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				logger.Error().Msg("Missing Authorization header")
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing Authorization header"})
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				logger.Error().Msg("Invalid Authorization header format")
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Authorization header format"})
			}

			userToken, err := security.VerifyToken(tokenString)
			if err != nil {
				logger.Error().Msgf("Failed to verify token: %v", err)
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			}

			userID, err := uuid.Parse(userToken.UserID)
			if err != nil {
				logger.Error().Msgf("Invalid user ID format: %v", err)
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
			}

			username := userToken.Username
			if username == "" {
				logger.Error().Msg("Invalid username")
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid username"})
			}

			c.Set("user_id", userID)
			c.Set("username", userToken.Username)

			logger.Info().Msgf("User ID: %s, Username: %s", userID, username)

			return next(c)
		}
	}
}
