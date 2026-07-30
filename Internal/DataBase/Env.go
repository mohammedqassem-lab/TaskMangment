package database

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func GetEnv() (Env, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Env{}, err
	}
	Env := Env{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbname:   os.Getenv("POSTGRES_DB"),
	}
	return Env, nil
}
func GetJwtKey() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	secrtKey := os.Getenv("jwtKey")
	return secrtKey, nil
}
