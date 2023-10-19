package domain

type ItemOrder struct {
	item     Item
	quantity int
}

func NewItemOrder(item Item, quantity int) *ItemOrder {
	return &ItemOrder{
		item:     item,
		quantity: quantity,
	}
}

func (io *ItemOrder) SetItem(item Item) {
	io.item = item
}

func (io *ItemOrder) SetQuantity(quantity int) {
	io.quantity = quantity
}

func (io *ItemOrder) ItemName() string {
	return io.item.name
}

func (io *ItemOrder) ItemId() int64 {
	return io.item.id
}

func (io *ItemOrder) ItemAmount() int {
	return io.item.amount
}

func (io *ItemOrder) Quantity() int {
	return io.quantity
}
