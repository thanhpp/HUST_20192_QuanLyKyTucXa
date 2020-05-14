package forms

//LoginForm receive login data
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
