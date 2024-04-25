package usecase

import "github.com/lourencodevopspython/api_mensageria/internal/entity"

type ListProductsOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type ListProductsUseCase struct {
	IProductRepository entity.IProductRepository
}

func NewListProductsUseCase(productRepository entity.IProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{IProductRepository: productRepository}
}

func (u *ListProductsUseCase) Execute() ([]*ListProductsOutputDto, error) {
	products, err := u.IProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var productsOutput []*ListProductsOutputDto
	for _, proproduct := range products {
		productsOutput = append(productsOutput, &ListProductsOutputDto{
			ID:    proproduct.ID,
			Name:  proproduct.Name,
			Price: proproduct.Price,
		})
	}
	return productsOutput, nil
}
