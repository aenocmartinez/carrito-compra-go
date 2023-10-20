package mysql

import (
	"bytes"
	"ejemplo2/src/domain"
	"ejemplo2/src/infraestructure/database"
	"ejemplo2/src/view/dto"
	"fmt"
)

type ItemDao struct {
}

func NewItemDao() *ItemDao {
	return &ItemDao{}
}

func (i *ItemDao) CreateItem(item domain.Item) error {

	db := database.InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("INSERT INTO items(name, amount) VALUES (?, ?)")

	stmt, err := db.Conn().Prepare(strSQL.String())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = stmt.Exec(item.Name(), item.Amount())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (i *ItemDao) DeleteItem(item domain.Item) error {
	return nil
}

func (i *ItemDao) FindItemById(id int64) domain.Item {
	return domain.Item{}
}

func (i *ItemDao) ItemList() (items []dto.ItemDto) {

	items = []dto.ItemDto{}
	var strQuery bytes.Buffer
	strQuery.WriteString("SELECT id, name, amount FROM items ORDER BY name")

	db := database.InstanceDB()
	rows, err := db.Conn().Query(strQuery.String())
	if err != nil {
		return items
	}
	defer rows.Close()

	for rows.Next() {
		var id, amount int
		var name string

		rows.Scan(&id, &name, &amount)

		items = append(items, dto.ItemDto{
			Id:     int64(id),
			Name:   name,
			Amount: amount,
		})
	}

	return items
}

func (i *ItemDao) FindItemByName(name string) domain.Item {
	db := database.InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT id, name, amount FROM items WHERE name=?")

	rs := db.Conn().QueryRow(strSQL.String(), name)

	var id, amount int
	var nameFound string

	rs.Scan(&id, &nameFound, &amount)

	item := domain.NewItem(nameFound, amount)
	item.SetId(int64(id))
	return *item
}
