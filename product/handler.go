package product

import (
	"context"
	"errors"
	"fmt"

	"encore.app/shared"
	"github.com/go-playground/validator/v10"

	"encore.app/infrastructure"
	user "encore.app/user"
	"encore.dev/beta/errs"
)

type repositoryI interface {
	GetProductBySKU(productSKU string) (*Product, error)
	GetProductByUUID(uuid string) (*Product, error)
	UpdateProduct(product *Product) (*Product, error)
	Save(data *Product) error
	Delete(uuid string) error
	GetRoleUser(email string) (*user.UserRole, error)
	GetAllProducts() ([]*Product, error)
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

//encore:api public method=POST path=/product/get
func (s *Service) GetProductBySku(ctx context.Context, data *ProductGetSKUDTO) (*ProductDTO, error) {
	cntvw, err := s.repository.GetRoleUser(data.Email)
	if err != nil {
		return nil, user.ErrUserAdminNotFound
	}
	if cntvw.Role != "admin" {
		return nil, errors.New("INSUFICIENT_PERMISIONS")
	}
	cntvwproduct, errproduct := s.repository.GetProductBySKU(data.SKU)
	if errproduct != nil {
		return nil, ErrProductNotFound
	}

	finalResponse := &Product{
		UUID:         cntvwproduct.UUID,
		SKU:          cntvwproduct.SKU,
		Name:         cntvwproduct.Name,
		Price:        cntvwproduct.Price,
		Brand:        cntvwproduct.Brand,
		QueryCounter: cntvwproduct.QueryCounter,
	}

	return toProductDTO(finalResponse), nil
}

//encore:api public method=POST path=/products
func (s *Service) GetAllProducts(ctx context.Context, dto *ProductsGetDTO) (*ProductsDTO, error) {

	fmt.Println("//////", dto.Email)
	cntvw, err := s.repository.GetRoleUser(dto.Email)
	if err != nil {
		return nil, user.ErrUserAdminNotFound
	}
	if cntvw.Role != "admin" {
		return nil, errors.New("INSUFICIENT_PERMISIONS")
	}

	products, error := s.repository.GetAllProducts()
	if error != nil {
		return nil, ErrProductNotFound
	}

	response := &ProductsDTO{
		Products: toProductDTOs(products),
	}

	return response, nil
}

//encore:api public method=POST path=/product
func (s *Service) SaveProduct(ctx context.Context, dto *ProductSaveDTO) error {
	err := s.validator.Validate(dto)
	if err != nil {
		return s.validator.ParseValidatorError(err)
	}

	context.Background()

	cntvw, err := s.repository.GetRoleUser(dto.AdminEmail)
	if err != nil {
		return user.ErrUserAdminNotFound
	}
	if cntvw.Role != "admin" {
		return errors.New("INSUFICIENT_PERMISIONS")
	}

	productToInsert := generateProductToSave(dto)

	err = s.repository.Save(productToInsert)

	if err != nil {
		return handleAPIErrors(err)
	}
	return nil
}

//encore:api public method=POST path=/product/delete
func (s *Service) DeleteProduct(ctx context.Context, data *ProductDeleteDTO) error {

	context.Background()

	cntvw, err := s.repository.GetRoleUser(data.AdminEmail)
	if err != nil {
		return user.ErrUserAdminNotFound
	}
	if cntvw.Role != "admin" {
		return errors.New("INSUFICIENT_PERMISIONS")
	}
	err = s.repository.Delete(data.UUID)
	if err != nil {
		return err
	}

	return nil
}

//encore:api public method=POST path=/product/update
func (s *Service) UpdateProduct(ctx context.Context, dto *ProductRequestUpdateDTO) error {
	err := s.validator.Validate(dto)
	if err != nil {
		return s.validator.ParseValidatorError(err)
	}

	cntvw, err := s.repository.GetRoleUser(dto.Email)
	if err != nil {
		return user.ErrUserAdminNotFound
	}

	if cntvw.Role != "admin" {
		return errors.New("INSUFICIENT_PERMISIONS")
	}

	context.Background()

	user, err := s.repository.GetProductByUUID(dto.UUIDToSearch)

	if err != nil {
		return &errs.Error{
			Code:    errs.NotFound,
			Message: "No product was found",
		}
	}
	if dto.Brand != "" {
		user.Brand = dto.Brand
	}
	if dto.Name != "" {
		user.Name = dto.Name
	}
	if dto.Price != 0 {
		user.Price = dto.Price
	}
	user, err = s.repository.UpdateProduct(user)

	if err != nil {
		return &errs.Error{
			Code:    errs.Internal,
			Message: "Internal server error",
		}
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
