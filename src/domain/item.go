package domain

type Item struct {
	id             int64
	name           string
	amount         int
	itemRepository ItemRepository
}

func NewItem(name string, amount int) *Item {
	return &Item{
		name:   name,
		amount: amount,
	}
}

func (i *Item) SetRepository(itemRepository ItemRepository) {
	i.itemRepository = itemRepository
}

func (i *Item) SetId(id int64) {
	i.id = id
}

func (i *Item) SetName(name string) {
	i.name = name
}

func (i *Item) SetAmount(amount int) {
	i.amount = amount
}

func (i *Item) Id() int64 {
	return i.id
}

func (i *Item) Name() string {
	return i.name
}

func (i *Item) Amount() int {
	return i.amount
}

func (i *Item) Exists() bool {
	return i.id > 0
}

func (i *Item) Create() error {
	return i.itemRepository.CreateItem(*i)
}
