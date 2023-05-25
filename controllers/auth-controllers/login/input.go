package loginAuth

type LoginInput struct{
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}