package domain

type ItemRepository interface {
	CreateItem(item Item) error
	DeleteItem(item Item) error
	FindItemById(id int64) Item
	ItemList() []Item
	FindItemByName(name string) Item
}
