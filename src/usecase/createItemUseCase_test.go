package usecase_test

import (
	"ejemplo2/src/dao/mock"
	"ejemplo2/src/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateItemUseCase(t *testing.T) {
	var expectedId int64 = 5
	name := "Item 5"
	amount := 500

	var itemRepository domain.ItemRepository = mock.NewItemDao()

	item := itemRepository.FindItemByName(name)
	if item.Exists() {
		t.Error("Ya existe")
		return
	}

	item.SetRepository(itemRepository)
	item.SetName(name)
	item.SetAmount(amount)

	item.Create()

	itemFound := itemRepository.FindItemById(5)
	assert.Equal(t, itemFound.Id(), expectedId)
}
