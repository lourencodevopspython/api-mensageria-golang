package usecase

import "github.com/lourencodevopspython/api_mensageria/internal/entity"

type CreateProductInputDto struct {
	Name  string  `json: "name"`
	Price float64 `json: "price"`
}

type CreateProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductUseCase struct {
	IProductRepository entity.IProductRepository
}

func NewCreateProductsUseCase(productRepository entity.IProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{IProductRepository: productRepository}
}

func (u *CreateProductUseCase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error) {
	product := entity.NewProduct(input.Name, input.Price)
	err := u.IProductRepository.Create(product)

	if err != nil {
		return nil, err
	}

	return &CreateProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil

}
