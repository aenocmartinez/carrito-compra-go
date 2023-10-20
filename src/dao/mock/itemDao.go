package mock

import (
	"ejemplo2/src/domain"
	"ejemplo2/src/view/dto"
	"strconv"
	"strings"
)

type ItemDao struct {
	items []domain.Item
}

func NewItemDao() *ItemDao {
	return &ItemDao{
		items: load(),
	}
}

func (i *ItemDao) CreateItem(item domain.Item) error {
	nextId := len(i.items) + 1
	item.SetId(int64(nextId))
	i.items = append(i.items, item)
	return nil
}

func (i *ItemDao) DeleteItem(item domain.Item) error {
	arrayItem := []domain.Item{}

	// foundItem := i.FindItemById(item.Id())
	// if !foundItem.Exists() {
	// 	return errors.New("item not found")
	// }

	for _, it := range i.items {
		if it.Id() == item.Id() {
			continue
		}
		arrayItem = append(arrayItem, it)
	}

	i.items = arrayItem

	return nil
}

func (i *ItemDao) FindItemById(id int64) domain.Item {
	for _, item := range i.items {
		if item.Id() == id {
			return item
		}
	}
	return domain.Item{}
}

// func (i *ItemDao) FindItemById(id int64) dto.ItemDto {
// 	for _, item := range i.items {
// 		if item.Id() == id {
// 			return dto.ItemDto{
// 				Id:     item.Id(),
// 				Name:   item.Name(),
// 				Amount: item.Amount(),
// 			}
// 		}
// 	}
// 	return dto.ItemDto{}
// }

// func (i *ItemDao) ItemList() []domain.Item {
// 	return i.items
// }

func (i *ItemDao) ItemList() []dto.ItemDto {

	items := []dto.ItemDto{}
	for _, item := range i.items {
		items = append(items, dto.ItemDto{
			Id:     item.Id(),
			Name:   item.Name(),
			Amount: item.Amount(),
		})
	}

	return items
}

func load() (items []domain.Item) {

	for i := 1; i <= 4; i++ {
		item := *domain.NewItem("Item "+strconv.Itoa(i), i*100)
		item.SetId(int64(i))
		items = append(items, item)
	}

	return items
}

func (i *ItemDao) FindItemByName(name string) domain.Item {
	for _, item := range i.items {
		if strings.EqualFold(item.Name(), name) {
			return item
		}
	}
	return domain.Item{}
}
