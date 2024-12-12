package requests

type CreateRole struct {
	Name   string `json:"name" validate:"required"`
	Code   string `json:"code" validate:"required"`
	Active *bool  `json:"active"`
}
