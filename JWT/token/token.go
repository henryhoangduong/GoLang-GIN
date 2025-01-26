package token

import (
	models "jwt/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	jWTPrivateToken = "SecrteTokenSecrteToken"
	ip              = "192.168.0.107"
)

func GenerateToken(claims *models.JwtClaims, expirationTime time.Time) (string, error) {
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().Unix()
	claims.Issuer = ip
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString([]byte(jWTPrivateToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString, origin string)(bool, *models.JwtClaims){
    claims := &models.JwtClaims{}
}
