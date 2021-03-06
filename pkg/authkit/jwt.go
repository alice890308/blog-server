package authkit

import (
	"context"
	"fmt"
	"time"

	"github.com/alice890308/blog-server/pkg/logkit"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type JWT interface {
	Generate(userID string) (string, error)
	Verify(accessToken string) (*Payload, error)
}

type JWTConfig struct {
	SecretKey     string        `long:"secretKey" env:"SECRETKEY" description:"jwt secret key" required:"true"`
	TokenDuration time.Duration `long:"timeDuration" env:"TIMEDURATION" description:"jwt token duration" default:"60h"`
}

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type Payload struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

func NewJWTManager(ctx context.Context, conf *JWTConfig) *JWTManager {
	_ = logkit.FromContext(ctx).With(
		zap.String("tokenDuration", conf.TokenDuration.String()),
	)

	return &JWTManager{
		secretKey:     conf.SecretKey,
		tokenDuration: conf.TokenDuration,
	}
}

func (j *JWTManager) Generate(userID string) (string, error) {
	claims := Payload{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.tokenDuration).Unix(),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// Verify verifies the access token string and return a user claim if the token is valid
func (j *JWTManager) Verify(accessToken string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&Payload{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(j.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*Payload)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
