package formrequest

type CreateItemFormRequest struct {
	Name   string `json:"name" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}
