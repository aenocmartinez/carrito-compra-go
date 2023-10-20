package usecase

import (
	"ejemplo2/src/dao/mock"
	"ejemplo2/src/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestItemListUseCase(t *testing.T){

	
	var expected int64 = 4
	var itemRepository domain.ItemRepository = mock.NewItemDao()

	items := itemRepository.ItemList()

	assert.Equal(t, len(items), int(expected))
}