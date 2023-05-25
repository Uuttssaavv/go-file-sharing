package register

type RegisterResponse struct {
	ID        uint `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Image     string `json:"image"`
	Token     string `json:"token"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
