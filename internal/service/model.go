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
	Balance     int
	IsActive    bool
	CreatedAt   *time.Time
}

func NewUser(in InputUser) User {
	timeNow := time.Now()

	return User{
		ID:          uuid.New(),
		Name:        in.Name,
		Age:         in.Age,
		PhoneNumber: in.PhoneNumber,
		Balance:     in.Balance,
		IsActive:    true,
		CreatedAt:   &timeNow,
	}
}

type InputUser struct {
	Name        string
	Age         int
	PhoneNumber *string
	Balance     int
}

type InputTransfer struct {
	From   uuid.UUID
	To     uuid.UUID
	Amount int
}
