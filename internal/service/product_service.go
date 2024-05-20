package service

import "fmt"

// ProductService 인터페이스 정의
type ProductService interface {
	GetProduct(id int) string
}

type productService struct{}

func NewProductService() ProductService {
	return &productService{}
}

func (p *productService) GetProduct(id int) string {
	return fmt.Sprintf("Product ID: %d", id)
}
