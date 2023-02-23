package user

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	CollectionName = "users"
)

var (
	ErrUserNotFound      = errors.New("USER_NOT_FOUND")
	ErrUserAdminNotFound = errors.New("INSUFICIENT_PERMISIONS")
)

type Repository struct {
	FirestoreClient *firestore.Client
}

func NewRepository(firestoreClient *firestore.Client) *Repository {
	return &Repository{
		FirestoreClient: firestoreClient,
	}
}
func (r *Repository) Save(data *User) error {
	ctx := context.Background()
	_, _, err := r.FirestoreClient.Collection(CollectionName).Add(ctx, data.toInterface())
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetRoleUser(email string) (*UserRole, error) {
	ctx := context.Background()

	var (
		userDAO *UserRoleDAO
	)
	const whereKey = "user_email"

	iter := r.FirestoreClient.Collection(CollectionName).Where(whereKey, "==", email).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, ErrUserAdminNotFound
	}
	if err = doc.DataTo(&userDAO); err != nil {
		return nil, ErrUserAdminNotFound
	}
	user := &UserRole{
		Role: userDAO.Role,
	}
	return user, nil
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	ctx := context.Background()

	var (
		userDAO *UserDAO
	)
	const whereKey = "user_email"

	iter := r.FirestoreClient.Collection(CollectionName).Where(whereKey, "==", email).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, ErrUserNotFound
	}
	if err = doc.DataTo(&userDAO); err != nil {
		return nil, ErrUserNotFound
	}
	user := &User{
		UUID:  userDAO.UUID,
		Name:  userDAO.Name,
		Email: userDAO.Email,
		Role:  userDAO.Role,
	}
	return user, nil
}

func (r *Repository) DeleteUser(email string) error {
	ctx := context.Background()

	const whereKey = "user_email"

	collection := r.FirestoreClient.Collection(CollectionName)
	iter := collection.Where(whereKey, "==", email).Documents(ctx)
	docRef, err := iter.Next()
	if err == iterator.Done {
		return ErrUserNotFound
	}
	if err != nil {
		return err
	}
	_, err = docRef.Ref.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateUser(user *User) (*User, error) {
	var userDAO *UserDAO

	ctx := context.Background()

	iter := r.FirestoreClient.Collection(CollectionName).Where("user_email", "==", user.Email).Documents(ctx)

	doc, err := iter.Next()

	if err != nil {
		return nil, ErrUserNotFound
	}

	toUpdate := user.toInterface()
	doc.Ref.Set(ctx, toUpdate)

	if err := doc.DataTo(&userDAO); err != nil {
		return nil, err
	}

	return toDomain(userDAO), nil
}

func (r *Repository) GetAllUsers() ([]*User, error) {
	ctx := context.Background()

	var (
		products   []*User
		product    *User
		productDAO *UserDAO
	)

	iter := r.FirestoreClient.Collection(CollectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		if err = doc.DataTo(&productDAO); err != nil {
			return nil, err
		}

		product = &User{
			UUID:  productDAO.UUID,
			Name:  productDAO.Name,
			Email: productDAO.Email,
			Role:  productDAO.Role,
		}

		products = append(products, product)
	}

	if products == nil {
		return nil, ErrUserNotFound
	}

	return products, nil
}

func (r *Repository) GetAllAdmins() ([]*User, error) {
	ctx := context.Background()

	var (
		admins  []*User
		user    *User
		userDAO *UserDAO
	)
	const whereKey = "user_role"
	const roleAdmin = "admin"

	iter := r.FirestoreClient.Collection(CollectionName).Where(whereKey, "==", roleAdmin).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err = doc.DataTo(&userDAO); err != nil {
			return nil, ErrUserNotFound
		}
		user = &User{
			UUID:  userDAO.UUID,
			Name:  userDAO.Name,
			Email: userDAO.Email,
			Role:  userDAO.Role,
		}
		admins = append(admins, user)
	}

	return admins, nil
}
