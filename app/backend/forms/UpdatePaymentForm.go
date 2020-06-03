package forms

type UpdatePaymentForm struct {
	MoneyManageID int    `json:"money_manage_id" binding:"required"`
	Status        string `json:"status" binding:"required"`
}
