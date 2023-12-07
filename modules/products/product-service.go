package products

import "context"

type IProductService interface {
	CreateProduct(ctx context.Context, product Product) (*Product, error)
	GetAllProduct(ctx context.Context) ([]Product, error)
	GetProductById(ctx context.Context, id uint) (*Product, error)
	UpdateProductById(ctx context.Context, id uint, product Product) (*Product, error)
	DeleteProductById(ctx context.Context, id uint) (any, error)
}

type productService struct {
	productRepository IProductRepository
}

func NewProductService(productRepository IProductRepository) *productService {
	return &productService{
		productRepository: productRepository,
	}
}

func (p *productService) CreateProduct(ctx context.Context, product Product) (*Product, error) {
	return p.productRepository.CreateProduct(ctx, product)
}

func (p *productService) GetAllProduct(ctx context.Context) ([]Product, error) {
	return p.productRepository.GetAllProduct(ctx)
}

func (p *productService) GetProductById(ctx context.Context, id uint) (*Product, error) {
	return p.productRepository.GetProductById(ctx, id)
}

func (p *productService) UpdateProductById(ctx context.Context, id uint, product Product) (*Product, error) {
	return p.productRepository.UpdateProductById(ctx, id, product)
}

func (p *productService) DeleteProductById(ctx context.Context, id uint) (any, error) {
	return p.productRepository.DeleteProductById(ctx, id)
}
