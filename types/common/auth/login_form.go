package auth

type LoginForm struct {
	Handle     string `json:"handle"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
}
