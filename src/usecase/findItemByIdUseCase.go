package usecase

import (
	"ejemplo2/src/dao/mock"
	"ejemplo2/src/domain"
	"ejemplo2/src/view/dto"
)

type FindItemByIdUseCase struct{}

func (uc *FindItemByIdUseCase) Execute(id int64) (response dto.ResponseDto) {

	var itemRepository domain.ItemRepository = mock.NewItemDao()

	item := itemRepository.FindItemById(id)

	if !item.Exists() {
		response.Code = "404"
		response.Data = "No se ha encontrado el item"
		return response
	}

	 itemDto := dto.ItemDto{
						Id:     item.Id(),
						Name:   item.Name(),
						Amount: item.Amount(),
					}
	response.Code = "200"
	response.Data = itemDto

	return response
}
