package user

type UserDAO struct {
	UUID     string `firestore:"UUID"`
	Name     string `firestore:"user_name"`
	Password string `firestore:"user_password"`
	Email    string `firestore:"user_email"`
	Role     string `firestore:"user_role"`
}

type UserRoleDAO struct {
	Role string `firestore:"user_role"`
}

type UserDTO struct {
	UUID     string `json:"UUID"`
	Name     string `json:"user_name"`
	Password string `json:"user_password"`
	Email    string `json:"user_email"`
	Role     string `json:"user_role"`
}

type UserCreateDTO struct {
	Name  string `json:"user_name"`
	Email string `json:"user_email"`
	Role  string `json:"user_role"`
}
type UserCreateRequestDTO struct {
	AdminEmail string `json:"user_email_request"`
	Name       string `json:"user_name"`
	Email      string `json:"user_email"`
	Role       string `json:"user_role"`
}

type UseRequestDTO struct {
	AdminEmail    string `json:"user_email_request"`
	EmailToSearch string `json:"user_email_to_search"`
}
type UseRequestDeleteDTO struct {
	AdminEmail    string `json:"user_email_request"`
	EmailToDelete string `json:"user_email_to_delete"`
}
type UseRequestUpdateDTO struct {
	AdminEmail    string `json:"user_email_request"`
	EmailToSearch string `json:"user_email_to_delete"`
	Name          string `json:"user_name_to_update"`
}

func toUserDTO(data *User) *UserDTO {
	return &UserDTO{
		UUID:  data.UUID,
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
	}
}
