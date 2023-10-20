package formrequest

type FindItemByIdFormRequest struct {
	Id int64 `json:"id" binding:"required"`
}