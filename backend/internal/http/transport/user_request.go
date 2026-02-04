package transport

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
