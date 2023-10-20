package domain

import "ejemplo2/src/view/dto"

type ItemRepository interface {
	CreateItem(item Item) error
	DeleteItem(item Item) error
	FindItemById(id int64) Item
	ItemList() []dto.ItemDto
	FindItemByName(name string) Item
}
