package domains

type User struct {
	Username string
	Password string
	UserID string
}

func NewUser(username_ string, password_ string, userID_ string) *User {
	return &User{
		Username: username_,
		Password: password_,
		UserID: userID_,
	}
}
