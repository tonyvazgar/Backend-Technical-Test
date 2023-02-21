package hello

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	collectionName = "test"
)

var (
	customErrorName        = errors.New("CUSTOM_ERROR_MESSAGE")
	ErrCounterViewNotFound = errors.New("COUNTER_VIEW_NOT_FOUND")
)

type repository struct {
	FirestoreClient *firestore.Client
}

func NewRepository(firestoreClient *firestore.Client) *repository {
	return &repository{
		FirestoreClient: firestoreClient,
	}
}

func (r *repository) Save(data *CounterView) error {
	ctx := context.Background()
	_, _, err := r.FirestoreClient.Collection(collectionName).Add(ctx, data.toInterface())
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByProductID(productId string) ([]*TestDAO, error) {
	ctx := context.Background()
	var (
		counterViews   []*TestDAO
		counterViewDAO *TestDAO
	)
	const whereKey = "hello"

	iter := r.FirestoreClient.Collection(collectionName).Where(whereKey, "==", productId).Documents(ctx)
	for {
		doc, err := iter.Next()
		fmt.Println("----->", doc, err)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, ErrCounterViewNotFound
		}
		if err = doc.DataTo(&counterViewDAO); err != nil {
			return nil, ErrCounterViewNotFound
		}
		counterViewDAO = &TestDAO{
			Hello: counterViewDAO.Hello,
		}
		counterViews = append(counterViews, counterViewDAO)
	}
	return counterViews, nil
}
