package service

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Name        string
	Age         int
	PhoneNumber *string
	IsActive    bool
	CreatedAt   *time.Time
}

func NewUser(in UserInput) User {
	timeNow := time.Now()

	return User{
		ID:          uuid.New(),
		Name:        in.Name,
		Age:         in.Age,
		PhoneNumber: in.PhoneNumber,
		IsActive:    true,
		CreatedAt:   &timeNow,
	}
}

type UserInput struct {
	Name        string
	Age         int
	PhoneNumber *string
}
