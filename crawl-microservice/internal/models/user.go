package models

import "crawl-microservice/internal/dto"

type Position string

const (
	PositionAdmin   Position = "ADMIN"
	PositionCourier Position = "COURIER"
	PositionUser    Position = "USER"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Position  string
	At
}

func NewUser(request dto.SignUpRequest) *User {
	return &User{
		FirstName: request.FisrtName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		Position:  "USER",
	}
}
