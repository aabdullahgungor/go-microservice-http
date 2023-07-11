package endpoint

import (
	"github.com/aabdullahgungor/go-microservice-http/product/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetAllProducts endpoint.Endpoint
	GetProductById endpoint.Endpoint
	CreateProduct  endpoint.Endpoint
	EditProduct    endpoint.Endpoint
	DeleteProduct  endpoint.Endpoint
}

func MakeEndpoints(s service.IProductService) Endpoints {
	return Endpoints{
		GetAllProducts: makeGetAllProductsEndpoint(s),
		GetProductById: makeGetProductByIdEndpoint(s),
		CreateProduct:  makeCreateProductEndpoint(s),
		EditProduct:    makeEditProductEndpoint(s),
		DeleteProduct:  makeDeleteProductEndpoint(s),
	}
}

func makeGetAllProductsEndpoint(s service.IProductService) endpoint.Endpoint {

}

func makeGetProductByIdEndpoint(s service.IProductService) endpoint.Endpoint {

}

func makeCreateProductEndpoint(s service.IProductService) endpoint.Endpoint {

}

func makeEditProductEndpoint(s service.IProductService) endpoint.Endpoint {

}

func makeDeleteProductEndpoint(s service.IProductService) endpoint.Endpoint {

}
