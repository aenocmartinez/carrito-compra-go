package controller

import (
	"ejemplo2/src/usecase"
	"ejemplo2/src/view/dto"
	formrequest "ejemplo2/src/view/form-request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var req formrequest.CreateItemFormRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemDto := dto.ItemDto{
		Name:   req.Name,
		Amount: req.Amount,
	}

	useCase := usecase.CreateItemUseCase{}
	response := useCase.Execute(itemDto)

	code, _ := strconv.Atoi(response.Code)
	c.JSON(code, response.Data)
}
