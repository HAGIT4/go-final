package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretKey = "notSoSecretKey"

type authClaims struct {
	jwt.StandardClaims
	Username string
}

type authService struct {
	secretKey string
}

func NewAuthService() (asv *authService) {
	return &authService{
		secretKey: secretKey,
	}
}

func (asv *authService) GenerateToken(username string) (token string, err error) {
	now := time.Now()
	jwtClaims := jwt.StandardClaims{
		ExpiresAt: now.Add(time.Hour * 24).Unix(),
		IssuedAt:  now.Unix(),
	}
	claims := &authClaims{
		jwtClaims,
		username,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (asv *authService) ValidateToken(encodedToken string) (token *jwt.Token, err error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("invalid token")
		}
		return []byte(asv.secretKey), nil
	}
	return jwt.Parse(encodedToken, keyFunc)
}
