package requests

type CreateProject struct {
	Name          string `json:"name" validate:"required"`
	Code          string `json:"code" validate:"required"`
	Connection_id string `json:"connection_id" validate:"required"`
}

type CreateDeviceProject struct {
	Device_id   string `json:"device_id" validate:"required"`
	Project_id  string `json:"project_id" validate:"required"`
	Activity_id string `json:"activity_id" validate:"required"`
}
