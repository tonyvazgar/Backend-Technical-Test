package user

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	UUID     string
	Name     string
	Password string
	Email    string
	Role     string
}

func (user *User) toInterface() map[string]interface{} {
	return map[string]interface{}{
		"UUID":          user.UUID,
		"user_name":     user.Name,
		"user_password": user.Password,
		"user_email":    user.Email,
		"user_role":     user.Role,
	}
}

func generateUserToSave(dto *UserDTO) *User {
	myUUID := uuid.New()
	counterViewToInsert := &User{
		UUID:     myUUID.String(),
		Name:     dto.Name,
		Password: dto.Password,
		Email:    dto.Email,
		Role:     dto.Role,
	}
	return counterViewToInsert
}

func newUserFromMap(data map[string]interface{}) *User {

	return &User{
		UUID:     fmt.Sprintf("%v", data["UUID"]),
		Name:     fmt.Sprintf("%v", data["user_name"]),
		Password: fmt.Sprintf("%v", data["user_password"]),
		Email:    fmt.Sprintf("%v", data["user_email"]),
		Role:     fmt.Sprintf("%v", data["user_role"]),
	}
}
