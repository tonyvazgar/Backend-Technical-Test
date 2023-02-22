package user

import (
	"context"
	"errors"

	"encore.app/infrastructure"
	"encore.app/shared"
	"encore.dev/beta/errs"
	"github.com/go-playground/validator/v10"
)

type repositoryI interface {
	Save(data *User) error
	GetRoleUser(email string) (*UserRole, error)
	GetUserByEmail(email string) (*User, error)
	DeleteUser(email string) error
	UpdateUser(user *User) (*User, error)
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

//encore:api public method=POST path=/users/create
func (s *Service) CreateUser(ctx context.Context, dto *UserCreateRequestDTO) error {
	err := s.validator.Validate(dto)
	if err != nil {
		return s.validator.ParseValidatorError(err)
	}
	if dto.Role != "admin" && dto.Role != "anonym" {
		return errors.New("Role invalid, only accepted: 'admin' or 'anonym'")
	}

	context.Background()

	cntvw, err := s.repository.GetRoleUser(dto.AdminEmail)
	if err != nil {
		return ErrUserAdminNotFound
	}

	if cntvw.Role != "admin" {
		return errors.New("INSUFICIENT_PERMISIONS")
	}

	dataToInsert := &UserCreateDTO{
		Name:  dto.Name,
		Email: dto.Email,
		Role:  dto.Role,
	}

	userToInsert := generateUserToSave(dataToInsert)

	err = s.repository.Save(userToInsert)

	if err != nil {
		return handleAPIErrors(err)
	}

	return nil
}

//encore:api public method=POST path=/users/get
func (s *Service) GetUser(ctx context.Context, dto *UseRequestDTO) (*UserDTO, error) {
	err := s.validator.Validate(dto)
	if err != nil {
		return nil, s.validator.ParseValidatorError(err)
	}

	context.Background()

	cntvw, err := s.repository.GetRoleUser(dto.AdminEmail)
	if err != nil {
		return nil, ErrUserAdminNotFound
	}

	if cntvw.Role != "admin" {
		return nil, errors.New("INSUFICIENT_PERMISIONS")
	}

	cntvwe, error := s.repository.GetUserByEmail(dto.EmailToSearch)
	if error != nil {
		return nil, handleAPIErrors(error)
	}

	return toUserDTO(cntvwe), nil
}

//encore:api public method=POST path=/users/delete
func (s *Service) DeleteUser(ctx context.Context, dto *UseRequestDeleteDTO) error {
	err := s.validator.Validate(dto)
	if err != nil {
		return s.validator.ParseValidatorError(err)
	}

	cntvw, err := s.repository.GetRoleUser(dto.AdminEmail)
	if err != nil {
		return ErrUserAdminNotFound
	}

	if cntvw.Role != "admin" {
		return errors.New("INSUFICIENT_PERMISIONS")
	}

	context.Background()

	error := s.repository.DeleteUser(dto.EmailToDelete)
	if error != nil {
		return handleAPIErrors(error)
	}

	return nil
}

//encore:api public method=POST path=/users/update
func (s *Service) UpdateUser(ctx context.Context, dto *UseRequestUpdateDTO) error {
	err := s.validator.Validate(dto)
	if err != nil {
		return s.validator.ParseValidatorError(err)
	}

	cntvw, err := s.repository.GetRoleUser(dto.AdminEmail)
	if err != nil {
		return ErrUserAdminNotFound
	}

	if cntvw.Role != "admin" {
		return errors.New("INSUFICIENT_PERMISIONS")
	}

	context.Background()

	user, err := s.repository.GetUserByEmail(dto.EmailToSearch)

	if err != nil {
		return &errs.Error{
			Code:    errs.NotFound,
			Message: "No user was found",
		}
	}
	user.Name = dto.Name
	user, err = s.repository.UpdateUser(user)

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
	case ErrUserNotFound:
		return &errs.Error{
			Code:    errs.NotFound,
			Message: err.Error(),
		}
	default:
		return err
	}
}
