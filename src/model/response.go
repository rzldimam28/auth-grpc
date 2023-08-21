package model

type UserResponse struct {
	ID       string
	Username string
	Email    string
	Password string
}

func (u *UserResponse) String() string {
	return "ID: " + u.ID + ", Email: " + u.Email + ", Username: " + u.Username
}
