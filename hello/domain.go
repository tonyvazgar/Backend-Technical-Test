package hello

func (c *CounterView) toInterface() map[string]interface{} {
	return map[string]interface{}{
		"store_id":   c.StoreId,
		"product_id": c.ProductId,
		"counter":    c.Counter,
		"date":       c.Date,
		"UUID":       c.UUID,
	}
}
