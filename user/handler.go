package user

import (
	"context"

	"encore.app/infrastructure"
	"encore.app/shared"
	"encore.dev/beta/errs"
	"github.com/go-playground/validator/v10"
)

type repositoryI interface {
	Save(data *User) error
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

//encore:api public method=POST path=/users
func (s *Service) SaveUser(ctx context.Context, dto *UserDTO) error {
	err := s.validator.Validate(dto)
	if err != nil {
		return s.validator.ParseValidatorError(err)
	}

	context.Background()

	counterViewToInsert := generateUserToSave(dto)

	err = s.repository.Save(counterViewToInsert)

	if err != nil {
		return handleAPIErrors(err)
	}

	return nil
}

func handleAPIErrors(err error) error {
	switch err {
	case ErrUserNotFound:
		return &errs.Error{
			Code:    errs.NotFound,
			Message: err.Error(),
		}
	default:
		return err
	}
}
