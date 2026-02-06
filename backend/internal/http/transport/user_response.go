package transport

type CreateUserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserID   int32  `json:"userID"`
}

type AuthenticateUserResponse struct {
	UserID        int32  `json:"userID"`
	Username      string `json:"username"`
	GenerateCount int32  `json:"generateCount"`
}
