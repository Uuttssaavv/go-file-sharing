package loginAuth

// LoginInput provides a way to define validation rules and
// customize the JSON representation of struct fields.
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
