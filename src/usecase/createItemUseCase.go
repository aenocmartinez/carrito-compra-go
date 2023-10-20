package usecase

import (
	"ejemplo2/src/dao/mock"
	"ejemplo2/src/domain"
	"ejemplo2/src/view/dto"
)

type CreateItemUseCase struct{}

func (uc *CreateItemUseCase) Execute(itemDto dto.ItemDto) (response dto.ResponseDto) {

	var itemRepository domain.ItemRepository = mock.NewItemDao()

	item := itemRepository.FindItemByName(itemDto.Name)
	if item.Exists() {
		response.Code = "202"
		response.Data = "Ha ocurrido un error en el sistema"
	}

	item.SetRepository(itemRepository)
	item.SetName(itemDto.Name)
	item.SetAmount(itemDto.Amount)

	err := item.Create()
	if err != nil {
		response.Code = "500"
		response.Data = "Ha ocurrido un error en el sistema"
	}

	response.Code = "201"
	response.Data = "success"

	return response
}
