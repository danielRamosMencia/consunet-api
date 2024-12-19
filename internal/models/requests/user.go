package requests

type CreateUser struct {
	Username        string  `json:"username" validate:"required"`
	Email           string  `json:"email" validate:"required,email"`
	Password        string  `json:"password" validate:"required"`
	Subscription_id *string `json:"subscription_id"`
}

type UpdateUserSubscription struct {
	Subscription_id *string `json:"subscription_id" validate:"required"`
}
