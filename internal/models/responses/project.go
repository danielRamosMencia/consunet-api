package responses

type UserProjects struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Collabs struct {
	Id         string `json:"id"`
	Collab_id  string `json:"collab_id"`
	Collab     string `json:"collab"`
	Email      string `json:"email"`
	Permission string `json:"permission"`
}

type ProjectDevices struct {
	Id              string  `json:"id"`
	Device          string  `json:"device"`
	Activity        string  `json:"activity"`
	Max_consumption float32 `json:"max_consumption"`
	Min_consumption float32 `json:"min_consumption"`
}
