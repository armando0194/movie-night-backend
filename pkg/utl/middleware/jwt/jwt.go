package jwt

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	apperr "github.com/armando0194/movie-night-backend/pkg/utl/error"
	"github.com/armando0194/movie-night-backend/pkg/utl/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// New generates new JWT variable necessery for auth middleware
func New(secret, algo string, d int) *Service {
	signingMethod := jwt.GetSigningMethod(algo)
	if signingMethod == nil {
		panic("invalid jwt signing method")
	}
	return &Service{
		key:      []byte(secret),
		duration: time.Duration(d) * time.Minute,
		algo:     signingMethod,
	}
}

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	// Secret key used for signing.
	key []byte

	// Duration for which the jwt token is valid.
	duration time.Duration

	// JWT signing algorithm
	algo jwt.SigningMethod
}

// MWFunc makes JWT implement the Middleware interface.
func (j *Service) MWFunc() gin.HandlerFunc {

	return func(c *gin.Context) {
		token, err := j.ParseToken(c)
		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		id := int(claims["id"].(float64))
		username := claims["u"].(string)
		email := claims["e"].(string)
		role := model.AccessRole(claims["r"].(float64))

		c.Set("id", id)
		c.Set("username", username)
		c.Set("email", email)
		c.Set("role", role)

		c.Next()
	}
}

// ParseToken parses token from Authorization header
func (j *Service) ParseToken(c *gin.Context) (*jwt.Token, error) {

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		return nil, apperr.Unauthorized
	}
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, apperr.Unauthorized
	}

	fmt.Println(parts[1])
	return jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		if j.algo != token.Method {
			return nil, apperr.Generic
		}
		return j.key, nil
	})

}

// GenerateToken generates new JWT token and populates it with user data
func (j *Service) GenerateToken(u *model.User) (string, string, error) {
	expire := time.Now().Add(j.duration)

	token := jwt.NewWithClaims((j.algo), jwt.MapClaims{
		"id":  u.ID,
		"u":   u.Username,
		"e":   u.Email,
		"r":   u.Role.AccessLevel,
		"exp": expire.Unix(),
	})

	tokenString, err := token.SignedString(j.key)

	return tokenString, expire.Format(time.RFC3339), err
}
