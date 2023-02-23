package product

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/firestore"
	user "encore.app/user"
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

func (r *repository) GetProductByUUID(uuid string) (*Product, error) {
	ctx := context.Background()
	var (
		productDAO *ProductDAO
		product    Product
	)
	const whereKey = "UUID"

	iter := r.FirestoreClient.Collection(collectionName).Where(whereKey, "==", uuid).Documents(ctx)
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

func (r *repository) GetAllProducts() ([]*Product, error) {
	ctx := context.Background()

	var (
		products   []*Product
		product    *Product
		productDAO *ProductDAO
	)
	iter := r.FirestoreClient.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		if err = doc.DataTo(&productDAO); err != nil {
			return nil, err
		}

		product = &Product{
			Name:         productDAO.Name,
			Price:        productDAO.Price,
			UUID:         productDAO.UUID,
			SKU:          productDAO.SKU,
			Brand:        productDAO.Brand,
			QueryCounter: productDAO.QueryCounter,
		}

		products = append(products, product)
	}

	if products == nil {
		return nil, ErrProductNotFound
	}

	return products, nil
}

func (r *repository) GetRoleUser(email string) (*user.UserRole, error) {
	ctx := context.Background()

	var (
		userDAO *user.UserRoleDAO
	)
	const whereKey = "user_email"

	iter := r.FirestoreClient.Collection(user.CollectionName).Where(whereKey, "==", email).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, user.ErrUserAdminNotFound
	}
	if err = doc.DataTo(&userDAO); err != nil {
		return nil, user.ErrUserAdminNotFound
	}
	user := &user.UserRole{
		Role: userDAO.Role,
	}
	return user, nil
}

func (r *repository) GetAllUsersAdmins() ([]*user.UserEmailName, error) {
	ctx := context.Background()

	var (
		products   []*user.UserEmailName
		product    *user.UserEmailName
		productDAO *user.UserDAO
	)
	whereKey := "user_role"
	iter := r.FirestoreClient.Collection(user.CollectionName).Where(whereKey, "==", "admin").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		if err = doc.DataTo(&productDAO); err != nil {
			return nil, err
		}

		fmt.Println("!!!!!", productDAO)
		product = &user.UserEmailName{
			Name:  productDAO.Name,
			Email: productDAO.Email,
		}
		fmt.Println("????", product)
		products = append(products, product)
	}

	if products == nil {
		return nil, user.ErrUserNotFound
	}

	return products, nil
}

func incrementCounterProduct(doc *firestore.DocumentSnapshot) {
	ctx := context.Background()
	_, _ = doc.Ref.Update(ctx, []firestore.Update{
		{Path: "query_counter", Value: firestore.Increment(1)},
	})
}

func (r *repository) UpdateProduct(user *Product) (*Product, error) {
	var userDAO *ProductDAO

	ctx := context.Background()

	iter := r.FirestoreClient.Collection(collectionName).Where("UUID", "==", user.UUID).Documents(ctx)

	doc, err := iter.Next()

	if err != nil {
		return nil, ErrProductNotFound
	}

	toUpdate := user.toInterface()
	doc.Ref.Set(ctx, toUpdate)

	if err := doc.DataTo(&userDAO); err != nil {
		return nil, err
	}

	return toDomain(userDAO), nil
}

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
