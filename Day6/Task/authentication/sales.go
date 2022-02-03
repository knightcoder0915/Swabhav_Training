package authentication

type Sales struct {
	Product  string
	Quantity int
}

func NewSale(product string, qty int) *Sales {
	return &Sales{Product: product, Quantity: qty}
}
