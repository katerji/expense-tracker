package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/katerji/expense-tracker/env"
	"os"
	"strconv"
	"time"
)

type JWTService struct{}

type customJWTClaims struct {
	User      User  `json:"user"`
	ExpiresAt int64 `json:"expires_at"`
}

func (jwtService JWTService) VerifyToken(token string) (User, error) {
	jwtSecret := env.JWTToken()
	return jwtService.validateToken(token, jwtSecret)
}

func (jwtService JWTService) VerifyRefreshToken(token string) (User, error) {
	jwtSecret := env.JWTRefreshToken()
	return jwtService.validateToken(token, jwtSecret)
}

func (jwtService JWTService) validateToken(token, jwtSecret string) (User, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return User{}, errors.New("error parsing token")
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		jsonClaims, err := json.Marshal(claims)
		if err != nil {
			return User{}, errors.New("error parsing token")
		}
		var customClaims customJWTClaims
		if err := json.Unmarshal(jsonClaims, &customClaims); err != nil {
			return User{}, errors.New("error parsing token")
		}
		expiresAt := time.Unix(customClaims.ExpiresAt, 0)
		if expiresAt.Before(time.Now()) {
			return User{}, errors.New("token expired")
		}
		return customClaims.User, nil
	}

	return User{}, errors.New("invalid token")
}

func (jwtService JWTService) CreateJwt(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":       user,
		"expires_at": getJWTExpiry(),
	})
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (jwtService JWTService) CreateRefreshJwt(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":       user,
		"expires_at": getJWTRefreshExpiry(),
	})
	jwtSecret := os.Getenv("JWT_REFRESH_SECRET")
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getJWTExpiry() int64 {
	expiryString := os.Getenv("JWT_EXPIRY")
	expiry, _ := strconv.Atoi(expiryString)
	return intToUnixTime(expiry)
}

func getJWTRefreshExpiry() int64 {
	expiryString := os.Getenv("JWT_REFRESH_EXPIRY")
	expiry, _ := strconv.Atoi(expiryString)
	return intToUnixTime(expiry)
}

func intToUnixTime(num int) int64 {
	now := time.Now()
	duration := time.Duration(num) * time.Second
	result := now.Add(duration)
	return result.Unix()
}
