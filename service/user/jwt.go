package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/katerji/expense-tracker/env"
	"time"
)

type customJWTClaims struct {
	User      *User `json:"user"`
	ExpiresAt int64 `json:"expires_at"`
}

func (s *Service) verifyToken(token string) (*User, error) {
	jwtSecret := env.JWTToken()
	return s.validateToken(token, jwtSecret)
}

func (s *Service) verifyRefreshToken(token string) (*User, error) {
	jwtSecret := env.JWTRefreshToken()
	return s.validateToken(token, jwtSecret)
}

func (s *Service) validateToken(token, jwtSecret string) (*User, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, errors.New("error parsing token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	jsonClaims, err := json.Marshal(claims)
	if err != nil {
		return nil, errors.New("error parsing token")
	}

	var customClaims customJWTClaims
	if err := json.Unmarshal(jsonClaims, &customClaims); err != nil {
		return nil, errors.New("error parsing token")
	}

	expiresAt := time.Unix(customClaims.ExpiresAt, 0)
	if expiresAt.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return customClaims.User, nil
}

type jwtPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *Service) createJWTPair(user *User) (*jwtPair, error) {
	accessToken, err := s.createJwt(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.createRefreshJwt(user)
	if err != nil {
		return nil, err
	}

	return &jwtPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) createJwt(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":       user,
		"expires_at": getJWTExpiry(),
	})

	jwtSecret := env.JWTToken()

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *Service) createRefreshJwt(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":       user,
		"expires_at": getJWTRefreshExpiry(),
	})
	jwtSecret := env.JWTRefreshToken()

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getJWTExpiry() int64 {
	const oneWeekInSeconds = 604800

	return intToUnixTime(oneWeekInSeconds)
}

func getJWTRefreshExpiry() int64 {
	const oneMonthInSeconds = 2592000

	return intToUnixTime(oneMonthInSeconds)
}

func intToUnixTime(num int) int64 {
	now := time.Now()
	duration := time.Duration(num) * time.Second
	result := now.Add(duration)
	return result.Unix()
}
