package service

import (
	"context"
	"errors"

	"github.com/aabdullahgungor/go-microservice-http/product/model"
	"github.com/aabdullahgungor/go-microservice-http/product/repository"
	"github.com/go-kit/log"
)

var (
	ErrIDIsNotValid    = errors.New("id is not valid")
	ErrNameIsNotEmpty  = errors.New("product title cannot be empty")
	ErrProductNotFound = errors.New("product not found")
)

type IProductService interface {
	GetAll(ctx context.Context) ([]model.Product, error)
	GetById(ctx context.Context, id string) (model.Product, error)
	Create(ctx context.Context, product *model.Product) error
	Edit(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, id string) error
}

type DefaultProductService struct {
	productRepo repository.IProductRepository
	logger      log.Logger
}

func NewDefaultProductService(pRepo repository.IProductRepository, logger log.Logger) *DefaultProductService {
	return &DefaultProductService{
		productRepo: pRepo,
		logger:      logger,
	}
}

func (d *DefaultProductService) GetAll(ctx context.Context) ([]model.Product, error) {
	return d.productRepo.GetAllProducts(ctx)
}

func (d *DefaultProductService) GetById(ctx context.Context, id string) (model.Product, error) {

	product, err := d.productRepo.GetProductById(ctx, id)

	if err != nil {
		return model.Product{}, repository.ErrProductNotFound
	}

	return product, nil
}

func (d *DefaultProductService) Create(ctx context.Context, product *model.Product) error {

	if product.Name == "" {
		return ErrNameIsNotEmpty
	}

	return d.productRepo.CreateProduct(ctx, product)
}

func (d *DefaultProductService) Edit(ctx context.Context, product *model.Product) error {

	if product.Id.String() == "" {
		return ErrIDIsNotValid
	}
	if product.Name == "" {
		return ErrNameIsNotEmpty
	}

	err := d.productRepo.EditProduct(ctx, product)

	if err != nil {
		return repository.ErrProductNotFound
	}

	return nil
}

func (d *DefaultProductService) Delete(ctx context.Context, id string) error {

	err := d.productRepo.DeleteProduct(ctx, id)

	if err != nil {
		return repository.ErrProductNotFound
	}

	return nil
}
