package user

func toDomain(u *UserDAO) *User {
	return &User{
		Email: u.Email,
		UUID:  u.UUID,
		Name:  u.Name,
		Role:  u.Role,
	}
}
