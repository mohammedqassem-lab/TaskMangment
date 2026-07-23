package model

import "time"

type User struct {
	Id           int64 
	Name         string
	Email        string
	Hashpassword string
	Created_at   time.Time
	Updated_at   time.Time
}