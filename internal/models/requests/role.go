package requests

import "time"

type CreateRole struct {
	Id         string    `json:"id" validate:"len=32"`
	Name       string    `json:"name" validate:"required"`
	Code       string    `json:"code" validate:"required"`
	Active     bool      `json:"active" validate:"required"`
	Created_at time.Time `json:"created_at" validate:"required"`
}
