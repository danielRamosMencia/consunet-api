package responses

type UserData struct {
	Id              string  `json:"id"`
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	Active          bool    `json:"active"`
	Subscription_id *string `json:"subscription_id"`
}
