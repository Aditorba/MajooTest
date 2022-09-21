package dto

type Users struct {
	Id       int64  `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}
