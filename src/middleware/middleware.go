package middleware

import (
	"cats-social/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	helper helpers.HelperInterface
}

type MiddlewareInterface interface {
	AuthMiddleware(c *gin.Context)
}

func NewMiddleware(helper helpers.HelperInterface) MiddlewareInterface {
	return &Middleware{helper}
}

func (m *Middleware) AuthMiddleware(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")

	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, "authorization header is missing")
		c.Abort()
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, "invalid token format")
		c.Abort()
		return
	}

	// Validate JWT
	claims, err := m.helper.ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusForbidden, "invalid token")
		c.Abort()
		return
	}

	c.Set("user_id", claims.Subject)
	c.Next()
}
