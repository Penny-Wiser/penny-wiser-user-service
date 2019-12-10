package model

type UserRegisterRequest struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (r UserRegisterRequest) Validate() bool {
	return true
}
