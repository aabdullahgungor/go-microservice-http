package product

import "context"

type GetAllProductsRequest struct {
	ctx context.Context
}
type GetAllProductsResponse struct {
	Products []Product `json:"products, omiempty"`
	Err      error     `json:"err,omitempty"`
}

func (r GetAllProductsResponse) error() error { return r.Err }

type GetProductByIdRequest struct {
	ID string
}
type GetProductByIdResponse struct {
	Products Product `json:"product, omiempty"`
	Err      error   `json:"err,omitempty"`
}

func (r GetProductByIdResponse) error() error { return r.Err }

type CreateProductRequest struct {
	Product Product
}
type CreateProductResponse struct {
	Err error `json:"err,omitempty"`
}

func (r CreateProductResponse) error() error { return r.Err }

type EditProductRequest struct {
	Product Product
}
type EditProductResponse struct {
	Err error `json:"err,omitempty"`
}

func (r EditProductResponse) error() error { return r.Err }

type DeleteProductRequest struct {
	ID string
}
type DeleteProductResponse struct {
	Err error `json:"err,omitempty"`
}

func (r DeleteProductResponse) error() error { return r.Err }
