package usecase

import (
	"ejemplo2/src/dao/mysql"
	"ejemplo2/src/domain"
	"ejemplo2/src/view/dto"
)

type ItemListUseCase struct{}

func (uc *ItemListUseCase) Execute() (response dto.ResponseDto) {

	var itemRepository domain.ItemRepository = mysql.NewItemDao()

	items := itemRepository.ItemList()

	response.Code = "200"
	response.Data = items

	return response
}
