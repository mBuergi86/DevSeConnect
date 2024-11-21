package security

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
)

type jwtCustomClaims struct {
	UserID          string          `json:"user_id"`
	Username        string          `json:"username"`
	Firstname       string          `json:"first_name"`
	Lastname        string          `json:"last_name"`
	Email           string          `json:"email"`
	Bio             string          `json:"bio"`
	Profile_Picture string          `json:"profile_picture"`
	ExpiresAt       jwt.NumericDate `json:"exp"`
	jwt.RegisteredClaims
}

func (c *jwtCustomClaims) Valid() error {
	if time.Now().After(c.ExpiresAt.Time) {
		return errors.New("token has expired")
	}
	return nil
}

var (
	jwtSecret = os.Getenv("JWT_SECRET")
	logger    = zerolog.New(os.Stdout).With().Timestamp().Logger()
)

func VerifyToken(tokenString string) (*jwtCustomClaims, error) {
	if jwtSecret == "" {
		logger.Error().Msg("JWT secret is not set")
		return nil, errors.New("JWT secret is not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse token")
		return nil, err
	}

	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok || !token.Valid {
		logger.Error().Msg("Invalid token claims or signature")
		return nil, jwt.ErrSignatureInvalid
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		logger.Error().Msg("Token has expired")
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
