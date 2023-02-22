package user

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	UUID  string
	Name  string
	Email string
	Role  string
}

type UserRole struct {
	Role string
}

func (user *User) toInterface() map[string]interface{} {
	return map[string]interface{}{
		"UUID":       user.UUID,
		"user_name":  user.Name,
		"user_email": user.Email,
		"user_role":  user.Role,
	}
}

func generateUserToSave(dto *UserCreateDTO) *User {
	myUUID := uuid.New()
	userToInsert := &User{
		UUID:  myUUID.String(),
		Name:  dto.Name,
		Email: dto.Email,
		Role:  dto.Role,
	}
	return userToInsert
}

func newUserFromMap(data map[string]interface{}) *User {

	return &User{
		UUID:  fmt.Sprintf("%v", data["UUID"]),
		Name:  fmt.Sprintf("%v", data["user_name"]),
		Email: fmt.Sprintf("%v", data["user_email"]),
		Role:  fmt.Sprintf("%v", data["user_role"]),
	}
}
