package request

type User struct {
	Username string `json:"username" validate:"required,min=3,max=11,lowercase"`
	Password string `json:"password" validate:"required,min=8,passwd"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required"`
}
