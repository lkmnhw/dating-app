package auth

import (
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	os.Setenv("JWT_SECRET_KEY", "my_secret_key")
	defer os.Unsetenv("JWT_SECRET_KEY")

	email := "test@example.com"
	token, err := GenerateToken(email)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)
	assert.Equal(t, email, claims.Email)
	assert.True(t, claims.ExpiresAt > time.Now().Unix())
}

func TestValidateToken(t *testing.T) {
	os.Setenv("JWT_SECRET_KEY", "my_secret_key")
	defer os.Unsetenv("JWT_SECRET_KEY")

	email := "test@example.com"
	token, _ := GenerateToken(email)

	claims, err := validateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, email, claims.Email)

	invalidToken := "invalid.token.string"
	claims, err = validateToken(invalidToken)
	assert.Error(t, err)
	assert.Nil(t, claims)

	expiredClaims := &Claims{
		Email: "expired@example.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-time.Minute).Unix(),
		},
	}
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
	expiredTokenString, _ := expiredToken.SignedString(jwtKey)

	claims, err = validateToken(expiredTokenString)
	assert.Error(t, err)
	assert.Nil(t, claims)
}
