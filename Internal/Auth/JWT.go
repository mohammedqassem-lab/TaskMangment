package auth

import (
	database "TaskMangment/Internal/DataBase"
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	jwtutil "github.com/kittipat1413/go-common/util/jwt"
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"uid"`
}

func CreateToken(userID int) (string, error) {
	ctx := context.Background()
	signingKey, err := database.GetJwtKey()
	if err != nil {
		log.Fatal(err.Error())
	}
	var signingKeyByte = []byte(signingKey)
	manager, err := jwtutil.NewJWTManager(jwtutil.HS256, signingKeyByte)
	if err != nil {
		return "", err
	}

	claims := &MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			Issuer:    "Issuer",
			Subject:   "Subject",
		},
		UserID: strconv.Itoa(userID),
	}

	token, err := manager.CreateToken(ctx, claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string) (*MyCustomClaims, error) {
	ctx := context.Background()
	key, err := database.GetJwtKey()
	if err != nil {
		return &MyCustomClaims{}, err
	}
	manager, err := jwtutil.NewJWTManager(jwtutil.HS256, []byte(key))
	if err != nil {
		return nil, err
	}

	parsedClaims := &MyCustomClaims{}

	err = manager.ParseAndValidateToken(ctx, token, parsedClaims)
	if err != nil {
		return nil, err
	}

	return parsedClaims, nil
}

func GenerateRefreshToken() (string, error) {

	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
