package token

import (
	"fmt"
	"github.com/TheVovchenskiy/banners/configs"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateAccesToken(userID uint, username string, role string) (accessToken string, err error) {
	claims := &Claims{
		UserId:   fmt.Sprint(userID),
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: configs.AccesTokenExpiresAt,
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(configs.JwtKey)
	if err != nil {
		return
	}

	return
}

// func GenerateRefreshToken(userID int, username string) (refreshToken string, err error) {
// 	refreshClaims := &Claims{
// 		UserId:   fmt.Sprint(userID),
// 		Username: username,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: RefreshTokenExpiresAt,
// 		},
// 	}

// 	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(configs.RefreshJwtKey)
// 	if err != nil {
// 		return
// 	}

// 	return
// }

func GenerateToken(userID uint, username string, role string) (accessToken string, err error) {
	accessToken, err = GenerateAccesToken(userID, username, role)
	if err != nil {
		return "", err
	}

	// refreshToken, err = GenerateRefreshToken(userID, username)
	// if err != nil {
	// 	return "", "", err
	// }

	return accessToken, nil
}
