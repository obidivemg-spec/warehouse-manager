package storage

import "fmt"

var Items = make(map[string]int)

func AddProduct(name string, amount int) {
	Items[name] += amount
}

func GetStock() map[string]int {
	return Items
}

func ReserveProduct(name string, amount int) error {
	value, exists := Items[name]
	if !exists || value < amount {
		return fmt.Errorf("Не хватает продукта на складе.\n")
	}
	newValue := value - amount
	Items[name] = newValue
	return nil
}
