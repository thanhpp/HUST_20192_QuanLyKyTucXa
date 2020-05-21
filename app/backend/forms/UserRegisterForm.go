package forms

//RegisterForm contains register information
type RegisterForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserID   int    `json:"userID" binding:"required"`
}
