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

func FindItemById(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	useCase := usecase.FindItemByIdUseCase{}
	response := useCase.Execute(id)

	code, _ := strconv.Atoi(response.Code)
	c.JSON(code, response.Data)
}

func ItemList(c *gin.Context) {

	useCase := usecase.ItemListUseCase{}
	response := useCase.Execute()

	c.JSON(http.StatusOK, response)
}
