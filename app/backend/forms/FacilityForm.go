package forms

type NewFacilityForm struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type NewFacilityManageForm struct {
	RoomID     int `json:"room_id" binding:"required"`
	FacilityID int `json:"facility_id" binding:"required"`
	Default    int `json:"default" binding:"required"`
}
