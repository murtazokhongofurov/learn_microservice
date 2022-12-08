package repo
import (
	pb "github.com/three_serviceSimple/product_service/genproto/product"
	pbp "github.com/three_serviceSimple/product_service/genproto/store"
)

type ProductStorageI interface {
	CreateProduct(*pbp.ProductRes) error
	GetProductById(*pb.ProductId) (*pb.Product, error)
	UpdateProduct(*pb.Product) (*pb.Product, error)
	DeleteProduct(*pb.ProductId) (*pb.Empty, error)
}
