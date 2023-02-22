package user

type UserDAO struct {
	UUID     string `firestore:"UUID"`
	Name     string `firestore:"user_name"`
	Password string `firestore:"user_password"`
	Email    string `firestore:"user_email"`
	Role     string `firestore:"user_role"`
}

type UserDTO struct {
	UUID     string `json:"UUID"`
	Name     string `json:"user_name"`
	Password string `json:"user_password"`
	Email    string `json:"user_email"`
	Role     string `json:"user_role"`
}

func toUserDTO(data *User) *UserDTO {
	return &UserDTO{
		UUID:     data.UUID,
		Name:     data.Name,
		Password: data.Password,
		Email:    data.Email,
		Role:     data.Role,
	}
}
