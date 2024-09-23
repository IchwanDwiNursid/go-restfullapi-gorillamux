package web

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=5"`
}

type UserResponse struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}

type LoginUserPayload struct {
	Email string `json:"email" validate:"required,email"`
}

