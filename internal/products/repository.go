package products

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    sellerID, //"FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	return prodList, nil
}
