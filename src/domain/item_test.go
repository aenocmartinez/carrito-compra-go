package domain_test

// import (
// 	"ejemplo2/src/dao/mock"
// 	"ejemplo2/src/domain"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestItemList(t *testing.T) {
// 	var expected int = 4
// 	var itemRepository domain.ItemRepository = mock.NewItemDao()

// 	items := itemRepository.ItemList()
// 	assert.Equal(t, len(items), expected)
// }

// func TestCreateItem(t *testing.T) {
// 	var expected int = 5

// 	item := domain.NewItem("Item 5", 500)

// 	var itemRepository domain.ItemRepository = mock.NewItemDao()
// 	itemRepository.CreateItem(*item)

// 	items := itemRepository.ItemList()
// 	assert.Equal(t, len(items), expected)
// }

// func TestDeleteItem(t *testing.T){
// 	var expected int = 3

// 	var itemRepository domain.ItemRepository = mock.NewItemDao()
// 	item := itemRepository.FindItemById(4)
// 	itemRepository.DeleteItem(item)

// 	items := itemRepository.ItemList()
// 	assert.Equal(t, len(items), expected) 
// }

// func TestFindItemById(t *testing.T){
// 	var expected int64 = 2
// 	var itemRepository domain.ItemRepository = mock.NewItemDao()
// 	item := itemRepository.FindItemById(2)
// 	assert.Equal(t, item.Id(), expected)
// }