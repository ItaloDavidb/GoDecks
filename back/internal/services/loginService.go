package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/italodavidb/goCrud/internal/models"
	"github.com/italodavidb/goCrud/internal/repository"
	"github.com/italodavidb/goCrud/internal/utils/hashUtils"
	"github.com/italodavidb/goCrud/internal/utils/jwtUtils"
)

func Login(user models.User) (string, error) {
	existingUser, err := repository.FetchUser(user.Username)
	if err != nil {
		return "", err
	}

	if err := hashUtils.CheckPasswordHash(user.Password, existingUser.Password); err != nil {
		return "", errors.New("senha incorreta")
	}

	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &jwtUtils.UserClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtUtils.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
