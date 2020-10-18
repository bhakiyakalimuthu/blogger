package svc

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	EmailID     string    `json:"emailId"`
	PhoneNumber string    `json:"phoneNumber"`
}

type Payload struct {
	Name        string `json:"name"`
	EmailID     string `json:"emailId"`
	PhoneNumber string `json:"phoneNumber"`
}

type StoreUserModel struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	EmailID     string    `db:"email_id"`
	PhoneNumber string    `db:"phone_number"`
}
