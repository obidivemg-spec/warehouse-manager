package main

import (
	"fmt"
	"warehouse-manager/storage"
)

func main() {
	var choose int
	var name string
	var count int

	for {
		fmt.Println("Выберите действие:\n 1. Добавить продукт\n 2. Состояние склада\n 3. Зарезервировать товар\n 4. Выйти")
		fmt.Scan(&choose)

		switch choose {
		case 1:
			fmt.Println("Введите название и количество:")
			fmt.Scan(&name, &count)
			storage.AddProduct(name, count)
			fmt.Println("Успешно добавлено!")
		case 2:
			stock := storage.GetStock()
			fmt.Println("--- Текущий склад ---")
			for item, amount := range stock {
				fmt.Printf("%s: %d\n", item, amount)
			}
			fmt.Println("---------------------")
		case 3:
			fmt.Println("Что резервируем и сколько?")
			fmt.Scan(&name, &count)
			err := storage.ReserveProduct(name, count)
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Println("Товар зарезервирован!")
			}
		case 4:
			fmt.Println("Завершение работы...")
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
