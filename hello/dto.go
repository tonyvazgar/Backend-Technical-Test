package hello

import "time"

type CounterViewDAO struct {
	StoreId   string    `firestore:"store_id"`
	ProductId string    `firestore:"product_id"`
	Counter   int64     `firestore:"counter"`
	Date      time.Time `firestore:"date"`
	UUID      string    `firestore:"UUID"`
}

type TestDAO struct {
	Hello string `firestore:"hello"`
}
