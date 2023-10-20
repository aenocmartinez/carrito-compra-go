package usecase

import (
	"ejemplo2/src/dao/mock"
	"ejemplo2/src/domain"
	"ejemplo2/src/view/dto"
	"fmt"
)

type ItemListUseCase struct {}

func (uc* ItemListUseCase) Execute() (response dto.ResponseDto){

	
	var itemRepository domain.ItemRepository = mock.NewItemDao()

	items := itemRepository.ItemList()
	fmt.Println(items)

	response.Code = "200"
	response.Data = items

	return response
}