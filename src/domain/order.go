package domain

import "fmt"

type Order struct {
	number int64
	items  []ItemOrder
}

func NewOrder(number int64) *Order {
	return &Order{
		number: number,
		items:  []ItemOrder{},
	}
}

func (o *Order) SetNumber(number int64) {
	o.number = number
}

func (o *Order) AddItem(item Item, quantity int) {
	o.items = append(o.items, *NewItemOrder(item, quantity))
}

func (o *Order) RemoveItem(item ItemOrder) {

}

func (o *Order) Items() []ItemOrder {
	return o.items
}

func (o *Order) Pay(method MethodPay) {
	fmt.Println("Orden de compra No. ", o.number)
	method.Pay(o.calculateTotal())
}

func (o *Order) calculateTotal() int {
	var total int = 0
	for _, item := range o.items {
		total = total + (item.ItemAmount() * item.Quantity())
	}

	return total
}
