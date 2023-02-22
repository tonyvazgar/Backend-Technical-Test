package user

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	collectionName = "users"
)

var (
	ErrUserNotFound = errors.New("USER_NOT_FOUND")
)

type repository struct {
	FirestoreClient *firestore.Client
}

func NewRepository(firestoreClient *firestore.Client) *repository {
	return &repository{
		FirestoreClient: firestoreClient,
	}
}
func (r *repository) Save(data *User) error {
	ctx := context.Background()
	_, _, err := r.FirestoreClient.Collection(collectionName).Add(ctx, data.toInterface())
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllAdmins() ([]*User, error) {
	ctx := context.Background()

	var (
		counterViews   []*User
		counterView    *User
		counterViewDAO *UserDAO
	)
	const whereKey = "user_role"
	const roleAdmin = "admin"

	iter := r.FirestoreClient.Collection(collectionName).Where(whereKey, "==", roleAdmin).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err = doc.DataTo(&counterViewDAO); err != nil {
			return nil, ErrUserNotFound
		}
		counterView = &User{
			UUID:     counterViewDAO.UUID,
			Name:     counterViewDAO.Name,
			Password: counterViewDAO.Password,
			Email:    counterViewDAO.Email,
			Role:     counterViewDAO.Role,
		}
		counterViews = append(counterViews, counterView)
	}

	return counterViews, nil
}
