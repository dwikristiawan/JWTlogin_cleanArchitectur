package request

type SignRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
