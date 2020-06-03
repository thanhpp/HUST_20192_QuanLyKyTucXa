package forms

//NewRequestForm ...
type NewRequestForm struct {
	Title   string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type ReplyRequestForm struct {
	RequestID int    `json:"request_id" binding:"required"`
	Reply     string `json:"reply" binding:"required"`
}
