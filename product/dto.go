package product

type ProductDAO struct {
	UUID         string `firestore:"UUID"`
	SKU          string `firestore:"product_sku"`
	Name         string `firestore:"product_name"`
	Price        int32  `firestore:"product_price"`
	Brand        string `firestore:"product_brand"`
	QueryCounter int32  `firestore:"query_counter"`
}

type ProductDTO struct {
	// UUID         string  `json:"UUID"`
	SKU          string `json:"product_sku"`
	Name         string `json:"product_name"`
	Price        int32  `json:"product_price"`
	Brand        string `json:"product_brand"`
	QueryCounter int32  `json:"query_counter"`
}

func toProductDTO(data *Product) *ProductDTO {
	return &ProductDTO{
		SKU:          data.SKU,
		Name:         data.Name,
		Price:        data.Price,
		Brand:        data.Brand,
		QueryCounter: data.QueryCounter,
	}
}
