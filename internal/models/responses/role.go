package responses

import "time"

type Role struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Active     bool      `json:"active"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
