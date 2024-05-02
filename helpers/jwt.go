package helpers

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Helpers struct{}

func NewHelper() HelperInterface {
	return &Helpers{}
}

// Claims structure to hold JWT claims
type Claims struct {
	jwt.StandardClaims
}

func (h *Helpers) GenerateToken(userID int) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_PRIVATE_KEY"))

	block, _ := pem.Decode(jwtSecret)
	if block == nil {
		return "", fmt.Errorf("failed to decode PEM block")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	claims := jwt.MapClaims{
		"sub": fmt.Sprintf("%v", userID),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(privateKey)
}

// ValidateJWT validates the JWT token
func (h *Helpers) ValidateJWT(tokenString string) (*Claims, error) {
	jwtSecret := []byte(os.Getenv("JWT_PUBLIC_KEY"))

	block, _ := pem.Decode(jwtSecret)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token validation failed: %v", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
