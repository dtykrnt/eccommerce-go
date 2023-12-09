package products

import (
	"context"
	"fmt"
	"golang-basic/models"

	"gorm.io/gorm"
)

type Products models.Products
type IProductRepository interface {
	CreateProduct(context.Context, Products) (*Products, error)
	GetAllProduct(context.Context) ([]Products, error)
	GetProductById(context.Context, uint) (*Products, error)
	UpdateProductById(context.Context, uint, Products) (*Products, error)
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

func (p *productRepository) CreateProduct(ctx context.Context, product Products) (*Products, error) {

	if err := p.db.WithContext(ctx).Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepository) GetAllProduct(ctx context.Context) ([]Products, error) {
	var products []Products
	if err := p.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}

	var result []Products
	for _, gw := range products {
		result = append(result, Products(gw))
	}

	return result, nil
}

func (p *productRepository) GetProductById(ctx context.Context, id uint) (*Products, error) {
	var product Products
	if err := p.db.WithContext(ctx).First(&product, id).Error; err != nil {
		return nil, err
	}
	result := Products(product)
	return &result, nil
}

func (p *productRepository) UpdateProductById(ctx context.Context, id uint, product Products) (*Products, error) {
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

	result := Products(product)
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
