package transport

type CreateUserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserID uint32 `json:"userID"`
}

type AuthenticateUserResponse struct {
	UserID uint32 `json:"userID"`
	Success bool `json:"success"`
}
