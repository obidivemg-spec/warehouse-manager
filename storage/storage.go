package storage

var Items = make(map[string]int)

func AddProduct(name string, amount int) {
	Items[name] += amount
}
