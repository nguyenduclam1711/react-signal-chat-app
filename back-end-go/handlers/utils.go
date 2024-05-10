package handlers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nguyenduclam1711/react-signal-chat-app/env"
	"github.com/nguyenduclam1711/react-signal-chat-app/models"
)

func GenerateAccessToken(user models.UserParseFromDB) (string, error) {
	claims := jwt.MapClaims{
		// 1 week exp
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(),
		"username": user.Username,
		"id":       user.Id,
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encoded token and send it as response
	return token.SignedString([]byte(env.EnvData["JWT_SECRET"]))
}

func GenerateRefreshToken(accessToken string) (string, error) {
	claims := jwt.MapClaims{
		// 2 weeks exp
		"exp":   time.Now().Add(14 * 24 * time.Hour).Unix(),
		"token": accessToken,
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encoded token and send it as response
	return token.SignedString([]byte(env.EnvData["JWT_SECRET"]))
}

func GenerateAuthTokens(user models.UserParseFromDB) (string, string, error) {
	accessToken, accessTokenErr := GenerateAccessToken(user)
	if accessTokenErr != nil {
		return "", "", accessTokenErr
	}
	refreshToken, refreshTokenErr := GenerateRefreshToken(accessToken)
	if refreshTokenErr != nil {
		return accessToken, "", refreshTokenErr
	}
	return accessToken, refreshToken, nil
}
