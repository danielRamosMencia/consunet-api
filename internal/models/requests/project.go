package requests

type CreateProject struct {
	Name          string `json:"name" validate:"required"`
	Code          string `json:"code" validate:"required"`
	Connection_id string `json:"conection_id" validate:"required"`
}
