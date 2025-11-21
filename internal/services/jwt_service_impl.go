package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtServiceImpl struct {
	secretKey []byte
}

func NewJWTService(secretKey string) JWTService {
	return &jwtServiceImpl{
		secretKey: []byte(secretKey),
	}
}

func (s *jwtServiceImpl) GenerateToken(userID, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *jwtServiceImpl) ValidateToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok1 := claims["user_id"].(string)
		email, ok2 := claims["email"].(string)

		if !ok1 || !ok2 {
			return nil, errors.New("invalid token claims")
		}

		return &TokenClaims{
			UserID: userID,
			Email:  email,
		}, nil
	}

	return nil, errors.New("invalid token")
}

