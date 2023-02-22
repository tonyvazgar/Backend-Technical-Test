package product

import (
	"fmt"

	"github.com/google/uuid"
)

type Product struct {
	UUID         string
	SKU          string
	Name         string
	Price        int32
	Brand        string
	QueryCounter int32
}

func (product *Product) toInterface() map[string]interface{} {
	return map[string]interface{}{
		"UUID":          product.UUID,
		"product_sku":   product.SKU,
		"product_name":  product.Name,
		"product_price": product.Price,
		"product_brand": product.Brand,
		"query_counter": product.QueryCounter,
	}
}
func generateProductToSave(dto *ProductDTO) *Product {
	counter := 0
	myUUID := uuid.New()
	counterViewToInsert := &Product{
		UUID:         myUUID.String(),
		SKU:          dto.SKU,
		Name:         dto.Name,
		Price:        int32(dto.Price),
		Brand:        dto.Brand,
		QueryCounter: int32(counter),
	}
	return counterViewToInsert
}
func newProductFromMap(data map[string]interface{}) *Product {

	return &Product{
		UUID:         fmt.Sprintf("%v", data["uuid"]),
		SKU:          fmt.Sprintf("%v", data["product_sku"]),
		Name:         fmt.Sprintf("%v", data["product_name"]),
		Price:        data["product_price"].(int32),
		Brand:        fmt.Sprintf("%v", data["product_brand"]),
		QueryCounter: data["query_counter"].(int32),
	}
}
