package product

import (
	"context"

	"encore.app/shared"
	"github.com/go-playground/validator/v10"

	"encore.app/infrastructure"
	"encore.dev/beta/errs"
)

type repositoryI interface {
	GetProductBySKU(productSKU string) (*Product, error)
	Save(data *Product) error
	Delete(uuid string) error
}
type apiValidator interface {
	Validate(i interface{}) error
	ParseValidatorError(err error) error
}

//encore:service
type Service struct {
	repository repositoryI
	validator  apiValidator
}

func initService() (*Service, error) {
	client, err := infrastructure.InitFirebase()
	if err != nil {
		return nil, err
	}

	validator := shared.NewAPIValidator(validator.New())
	repository := NewRepository(client)

	return &Service{
		repository,
		validator,
	}, nil
}

//encore:api public method=GET path=/product/:sku
func (s *Service) GetProductSku(ctx context.Context, sku string) (*ProductDTO, error) {

	cntvw, err := s.repository.GetProductBySKU(sku)
	if err != nil {
		return nil, ErrProductNotFound
	}

	finalResponse := &Product{
		UUID:         cntvw.UUID,
		SKU:          cntvw.SKU,
		Name:         cntvw.Name,
		Price:        cntvw.Price,
		Brand:        cntvw.Brand,
		QueryCounter: cntvw.QueryCounter,
	}

	return toProductDTO(finalResponse), nil
}

//encore:api public method=POST path=/product
func (s *Service) SaveProduct(ctx context.Context, dto *ProductDTO) error {
	err := s.validator.Validate(dto)
	if err != nil {
		return s.validator.ParseValidatorError(err)
	}

	context.Background()

	productToInsert := generateProductToSave(dto)

	err = s.repository.Save(productToInsert)

	if err != nil {
		return handleAPIErrors(err)
	}

	// responseFoundProduct := &GetCounterViewResponseDTO{
	// 	CounterView: toCounterViewDTO(cntvw),
	// }
	// _, err = s.repository.UpdateViewsProduct(responseFoundProduct.CounterView.UUID)
	// if err != nil {
	// 	return handleAPIErrors(err)
	// }

	return nil
}

//encore:api public method=POST path=/product/delete
func (s *Service) DeleteProduct(ctx context.Context, data *ProductDeleteDTO) error {

	context.Background()

	err := s.repository.Delete(data.UUID)
	if err != nil {
		return err
	}

	return nil
}

func handleAPIErrors(err error) error {
	switch err {
	case ErrProductNotFound:
		return &errs.Error{
			Code:    errs.NotFound,
			Message: err.Error(),
		}
	default:
		return err
	}
}
