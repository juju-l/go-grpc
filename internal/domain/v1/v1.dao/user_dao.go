package v1_dao

type UserLoginReq struct {
	User string `json:"user"`
	Pswd string `json:"pswd"`
}

type User struct {
	User string `json:"user"`
}
