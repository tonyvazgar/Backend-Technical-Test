package user

type UserDAO struct {
	UUID  string `firestore:"UUID"`
	Name  string `firestore:"user_name"`
	Email string `firestore:"user_email"`
	Role  string `firestore:"user_role"`
}

type UserRoleDAO struct {
	Role string `firestore:"user_role"`
}

type UserDTO struct {
	UUID  string `json:"UUID"`
	Name  string `json:"user_name"`
	Email string `json:"user_email"`
	Role  string `json:"user_role"`
}
type UsersDTO struct {
	Users []*UserDTO `json:"users"`
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
type UsersGetDTO struct {
	Email string `json:"user_email_request"`
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
	EmailToSearch string `json:"user_email_to_search"`
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

func toUsersDTOs(data []*User) []*UserDTO {
	var arr []*UserDTO

	for lIndex := 0; lIndex < len(data); lIndex++ {
		arr = append(arr, newUserDTO(data[lIndex]))
	}

	return arr
}

func newUserDTO(data *User) *UserDTO {
	return &UserDTO{
		UUID:  data.UUID,
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
	}
}
