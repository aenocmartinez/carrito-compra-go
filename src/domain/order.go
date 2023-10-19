package domain

type Order struct {
	number int64
	items  []ItemOrder
}

func (o *Order) SetNumber(number int64) {
	o.number = number
}

func (o *Order) AddItem(item ItemOrder) {
	o.items = append(o.items, item)
}

func (o *Order) RemoveItem(item ItemOrder) {

}

func (o *Order) Items() []ItemOrder {
	return o.items
}

func (o *Order) Pay() {
	o.calculateTotal()
}

func (o *Order) calculateTotal() int {
	var total int = 0
	for _, item := range o.items {
		total = total + (item.ItemAmount() * item.Quantity())
	}

	return total
}
