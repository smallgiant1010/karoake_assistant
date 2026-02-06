package domains

type User struct {
	UserID int32
	Username string
	Password string
	GenerateCount int32
}

func NewUser(userID int32, username string, password string, generateCount int32) *User {
	return &User{
		UserID: userID,
		Username: username,
		Password: password,
		GenerateCount: generateCount,
	}
}
