package request

type UserRegister struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (r UserRegister) Validate() bool {
	return true
}
