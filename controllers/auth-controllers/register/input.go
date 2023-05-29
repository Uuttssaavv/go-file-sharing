package register

type RegisterInput struct {
	Username string `json:"username" validate:"required,lowercase"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	Image    string `json:"image"`
}
