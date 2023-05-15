package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenGenerator struct {
	secretKey string
}

type Claims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type Params struct {
	Id   string
	Role string
	Ttl  time.Duration
}

func NewTokenGenerator(cfg Config) *TokenGenerator {
	return &TokenGenerator{
		secretKey: cfg.SecretKey,
	}
}

func (g *TokenGenerator) GenerateNewAccessToken(params Params) (string, error) {
	claims := &Claims{
		Id:   params.Id,
		Role: params.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(params.Ttl).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(g.secretKey))
}

func (g *TokenGenerator) GenerateExpiredToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"exp": time.Now()})

	return token.SignedString([]byte(g.secretKey))
}
