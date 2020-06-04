package forms

type StudentInfoForm struct {
	Name    string `json:"name"`
	DOB     string `json:"DOB"`
	Contact string `json:"contact"`
	Address string `json:"address"`
}
