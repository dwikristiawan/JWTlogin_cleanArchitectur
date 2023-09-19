package model

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
