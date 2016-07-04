package order

import (
	"fmt"

	"../../database"
)

func (order *Order) Save() error {
	if !database.Storage.Opened {
		return fmt.Errorf("Database Error: DB must be opened before deleting.")
	}
	return database.Storage.DB.Save(order)
}

func (order Order) Delete() error {
	if !database.Storage.Opened {
		return fmt.Errorf("Database Error: DB must be opened before deleting.")
	}
	return database.Storage.DB.Remove(&order)
}

func (order Order) Get(key int) (Order, error) {
	if !database.Storage.Opened {
		return order, fmt.Errorf("Database Error: DB must be opened before deleting.")
	}
	err := database.Storage.DB.One("ID", key, &order)
	return order, err
}

// All returns all the orders
func All() ([]Order, error) {
	var err error
	var orders []Order
	if !database.Storage.Opened {
		return orders, fmt.Errorf("Database must be opened first.")
	}
	database.Storage.DB.All(&orders)
	return orders, err
}
