package product

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	collectionName = "products"
)

var (
	ErrProductNotFound = errors.New("PRODUCT_NOT_FOUND")
)

type repository struct {
	FirestoreClient *firestore.Client
}

func NewRepository(firestoreClient *firestore.Client) *repository {
	return &repository{
		FirestoreClient: firestoreClient,
	}
}

func (r *repository) Save(data *Product) error {
	ctx := context.Background()
	_, _, err := r.FirestoreClient.Collection(collectionName).Add(ctx, data.toInterface())
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetProductBySKU(productSKU string) (*Product, error) {
	ctx := context.Background()
	var (
		productDAO *ProductDAO
		product    Product
	)
	const whereKey = "product_sku"

	iter := r.FirestoreClient.Collection(collectionName).Where(whereKey, "==", productSKU).Documents(ctx)
	doc, err := iter.Next()
	if err != nil {
		return nil, ErrProductNotFound
	}

	incrementCounterProduct(doc)

	if err = doc.DataTo(&productDAO); err != nil {
		return nil, err
	}

	product = Product{
		UUID:         productDAO.UUID,
		SKU:          productDAO.SKU,
		Name:         productDAO.Name,
		Price:        productDAO.Price,
		Brand:        productDAO.Brand,
		QueryCounter: productDAO.QueryCounter,
	}

	return &product, nil
}

func incrementCounterProduct(doc *firestore.DocumentSnapshot) {
	ctx := context.Background()
	_, _ = doc.Ref.Update(ctx, []firestore.Update{
		{Path: "query_counter", Value: firestore.Increment(1)},
	})
}

// func (r *repository) UpdateViewsProduct(productID string) (*Product, error) {
// 	var cw *Product
// 	const whereKey = "UUID"
// 	ctx := context.Background()

// 	iter := r.FirestoreClient.Collection(collectionName).Where(whereKey, "==", productID).Documents(ctx)
// 	for {
// 		docs, err := iter.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			return nil, err
// 		}

// 		data := docs.Data()
// 		cw = newCounterViewFromMap(data)

// 		_, err = docs.Ref.Update(ctx, []firestore.Update{
// 			{Path: "counter", Value: firestore.Increment(1)},
// 		})
// 	}

// 	if cw == nil {
// 		return nil, ErrCounterViewNotFound
// 	}

// 	return cw, nil
// }

func (f *repository) Delete(uuid string) error {

	ctx := context.Background()
	collection := f.FirestoreClient.Collection(collectionName)
	iter := collection.Where("UUID", "==", uuid).Documents(ctx)
	docRef, err := iter.Next()
	if err == iterator.Done {
		return ErrProductNotFound
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
