package model

import "time"

type User struct {
	UserUUID     string
	Email        string
	Name         string
	FirebaseUUID string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
