package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository}
}

// GetProduct returns a product by its ID
// This Product was not supposed to be returned. We should return a ProductDTO instead.
// However, we are returning a Product to simplify the example.
func (u *ProductUseCase) GetProduct(id int) (*Product, error) {
	return u.repository.GetProduct(id)
}
