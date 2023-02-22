package hello

import (
	"context"
	"fmt"
	"time"

	// "encore.app/shared"

	"encore.app/infrastructure"
	"encore.dev/beta/errs"
)

type CounterView struct {
	StoreId   string
	ProductId string
	Counter   int64
	Date      time.Time
	UUID      string
}

type repositoryI interface {
	Save(data *CounterView) error
	GetByProductID(productId string) ([]*TestDAO, error)
}
type apiValidator interface {
	Validate(i interface{}) error
	ParseValidatorError(err error) error
}

//encore:service
type Service struct {
	repository repositoryI
	// validator  apiValidator
}

func initService() (*Service, error) {
	client, err := infrastructure.InitFirebase()
	if err != nil {
		return nil, err
	}

	// validator := shared.NewAPIValidator(validator.New())
	repository := NewRepository(client)

	return &Service{
		repository,
		// validator,
	}, nil
}

type ViewProductDTO struct {
	Value string
}

type createCounterViewStoreDTO struct {
}

//encore:api public method=GET path=/test/:productId
func (s *Service) GetViewsFromProduct(ctx context.Context, productId string) (*TestDAO, error) {

	cntvw, err := s.repository.GetByProductID(productId)

	fmt.Println("Reponse is : ", cntvw, err)
	if err != nil {
		return nil, nil
	}
	if cntvw == nil {
		return nil, ErrCounterViewNotFound
	}

	return cntvw[len(cntvw)-1], nil
}

func handleAPIErrors(err error) error {
	switch err {
	case ErrCounterViewNotFound:
		return &errs.Error{
			Code:    errs.NotFound,
			Message: err.Error(),
		}
	default:
		return err
	}
}
