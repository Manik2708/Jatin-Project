package schemas

type CustomerLoginInput struct {
	Identify string `json:"identify"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type CustomerLoginOutput struct {
	BaseUser
	Token string `json:"token"`
}