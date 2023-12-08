package products

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(context.Context, Product) (*Product, error)
	GetAllProduct(context.Context) ([]Product, error)
	GetProductById(context.Context, uint) (*Product, error)
	UpdateProductById(context.Context, uint, Product) (*Product, error)
	DeleteProductById(context.Context, uint) (any, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) CreateProduct(ctx context.Context, product Product) (*Product, error) {
	newProduct := Product{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		IsActive:    product.IsActive,
	}
	if err := p.db.WithContext(ctx).Create(&product).Error; err != nil {
		return nil, err
	}
	result := Product(newProduct)

	fmt.Println(result)
	return &result, nil
}

func (p *productRepository) GetAllProduct(ctx context.Context) ([]Product, error) {
	var products []Product
	if err := p.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}

	var result []Product
	for _, gw := range products {
		result = append(result, Product(gw))
	}

	return result, nil
}

func (p *productRepository) GetProductById(ctx context.Context, id uint) (*Product, error) {
	var product Product
	if err := p.db.WithContext(ctx).First(&product, id).Error; err != nil {
		return nil, err
	}

	result := Product(product)

	fmt.Println(result)
	return &result, nil
}

func (p *productRepository) UpdateProductById(ctx context.Context, id uint, product Product) (*Product, error) {
	existingProduct, err := p.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}

	existingProduct.Name = product.Name
	existingProduct.Price = product.Price
	existingProduct.Description = product.Description
	existingProduct.IsActive = product.IsActive

	if err := p.db.WithContext(ctx).Save(&existingProduct).Error; err != nil {
		return nil, err
	}

	result := Product(product)
	result.ID = existingProduct.ID

	fmt.Println(result)
	return &result, nil
}

func (p *productRepository) DeleteProductById(ctx context.Context, id uint) (any, error) {

	existingProduct, err := p.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := p.db.WithContext(ctx).Delete(&existingProduct).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
