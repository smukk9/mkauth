package user

type User struct {
	Id       string
	Username string
	Email    string
}

type RequestUserBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResponseUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
