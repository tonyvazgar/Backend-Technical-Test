package product

type ProductDAO struct {
	UUID         string `firestore:"UUID"`
	SKU          string `firestore:"product_sku"`
	Name         string `firestore:"product_name"`
	Price        int32  `firestore:"product_price"`
	Brand        string `firestore:"product_brand"`
	QueryCounter int32  `firestore:"query_counter"`
}

func toDomain(u *ProductDAO) *Product {
	return &Product{
		UUID:         u.UUID,
		Name:         u.Name,
		SKU:          u.SKU,
		Price:        u.Price,
		Brand:        u.Brand,
		QueryCounter: u.QueryCounter,
	}
}
