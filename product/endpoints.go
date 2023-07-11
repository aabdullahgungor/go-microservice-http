package product

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetAllProducts endpoint.Endpoint
	GetProductById endpoint.Endpoint
	CreateProduct  endpoint.Endpoint
	EditProduct    endpoint.Endpoint
	DeleteProduct  endpoint.Endpoint
}

func MakeEndpoints(s IProductService) Endpoints {
	return Endpoints{
		GetAllProducts: makeGetAllProductsEndpoint(s),
		GetProductById: makeGetProductByIdEndpoint(s),
		CreateProduct:  makeCreateProductEndpoint(s),
		EditProduct:    makeEditProductEndpoint(s),
		DeleteProduct:  makeDeleteProductEndpoint(s),
	}
}

func makeGetAllProductsEndpoint(s IProductService) endpoint.Endpoint {

}

func makeGetProductByIdEndpoint(s IProductService) endpoint.Endpoint {

}

func makeCreateProductEndpoint(s IProductService) endpoint.Endpoint {

}

func makeEditProductEndpoint(s IProductService) endpoint.Endpoint {

}

func makeDeleteProductEndpoint(s IProductService) endpoint.Endpoint {

}
