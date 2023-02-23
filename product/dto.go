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
	AdminEmail   string `json:"user_email_request"`
	UUID         string `json:"UUID"`
	SKU          string `json:"product_sku"`
	Name         string `json:"product_name"`
	Price        int32  `json:"product_price"`
	Brand        string `json:"product_brand"`
	QueryCounter int32  `json:"query_counter"`
}
type ProductoDTO struct {
	UUID         string `json:"UUID"`
	SKU          string `json:"product_sku"`
	Name         string `json:"product_name"`
	Price        int32  `json:"product_price"`
	Brand        string `json:"product_brand"`
	QueryCounter int32  `json:"query_counter"`
}

type ProductsDTO struct {
	Products []*ProductoDTO `json:"products"`
}

type ProductSaveDTO struct {
	AdminEmail string `json:"user_email_request"`
	SKU        string `json:"product_sku"`
	Name       string `json:"product_name"`
	Price      int32  `json:"product_price"`
	Brand      string `json:"product_brand"`
}
type ProductsGetDTO struct {
	Email string `json:"user_email_request"`
}

type ProductDeleteDTO struct {
	AdminEmail string `json:"user_email_request"`
	UUID       string `json:"uuid"`
}

func toProductDTO(data *Product) *ProductDTO {
	return &ProductDTO{
		UUID:         data.UUID,
		SKU:          data.SKU,
		Name:         data.Name,
		Price:        data.Price,
		Brand:        data.Brand,
		QueryCounter: data.QueryCounter,
	}
}

func toProductDTOs(data []*Product) []*ProductoDTO {
	var arr []*ProductoDTO

	for lIndex := 0; lIndex < len(data); lIndex++ {
		arr = append(arr, newLegacyProductDTO(data[lIndex]))
	}

	return arr
}
func newLegacyProductDTO(data *Product) *ProductoDTO {
	return &ProductoDTO{
		UUID:         data.UUID,
		SKU:          data.SKU,
		Name:         data.Name,
		Price:        data.Price,
		Brand:        data.Brand,
		QueryCounter: data.QueryCounter,
	}
}
