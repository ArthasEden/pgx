package api

import (
	"pgxPractice/internal/service"

	"github.com/google/uuid"
)

type DTOreq struct {
	Name        string
	Age         int
	PhoneNumber *string
	Balance     int
}

type DTOresp struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	PhoneNumber *string   `json:"phone_number"`
	Balance     int
}

type DTOreqTransfer struct {
	From   uuid.UUID
	To     uuid.UUID
	Amount int
}

func DTOServiceToApi(users []service.User) []DTOresp {
	dtoUsers := make([]DTOresp, 0)

	for _, u := range users {
		dtoUser := DTOresp{
			ID:          u.ID,
			Name:        u.Name,
			Age:         u.Age,
			PhoneNumber: u.PhoneNumber,
			Balance:     u.Balance,
		}
		dtoUsers = append(dtoUsers, dtoUser)
	}

	return dtoUsers
}
