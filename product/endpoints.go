package product

import (
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Endpoints struct {
	GetAllProductsEndpoint endpoint.Endpoint
	GetProductByIdEndpoint endpoint.Endpoint
	CreateProductEndpoint  endpoint.Endpoint
	EditProductEndpoint    endpoint.Endpoint
	DeleteProductEndpoint  endpoint.Endpoint
}

func MakeServerEndpoints(s IProductService) Endpoints {
	return Endpoints{
		GetAllProductsEndpoint: makeGetAllProductsEndpoint(s),
		GetProductByIdEndpoint: makeGetProductByIdEndpoint(s),
		CreateProductEndpoint:  makeCreateProductEndpoint(s),
		EditProductEndpoint:    makeEditProductEndpoint(s),
		DeleteProductEndpoint:  makeDeleteProductEndpoint(s),
	}
}

func MakeClientEndpoints(instance string) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""

	options := []httptransport.ClientOption{}

	return Endpoints{
		GetAllProductsEndpoint: httptransport.NewClient("GET", tgt, decodeGetAllProductsRequest, encodeGetAllProductsResponse, options...).Endpoint(),
		GetProductByIdEndpoint: httptransport.NewClient("GET", tgt, decodeGetProductByIdRequest, encodeGetProductByIdResponse, options...).Endpoint(),
		CreateProductEndpoint:  httptransport.NewClient("POST", tgt, decodeCreateProductRequest, encodeCreateProductResponse, options...).Endpoint(),
		EditProductEndpoint:    httptransport.NewClient("PUT", tgt, decodeEditProductRequest, encodeEditProductResponse, options...).Endpoint(),
		DeleteProductEndpoint:  httptransport.NewClient("DELETE", tgt, decodeDeleteProductRequest, encodeDeleteProductResponse, options...).Endpoint(),
	}, nil
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
