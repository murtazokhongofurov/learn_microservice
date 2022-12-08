package service
import (
	"github.com/jmoiron/sqlx"
	l "github.com/three_serviceSimple/product_service/pkg/logger"
	"github.com/three_serviceSimple/product_service/storage"
)

type ProductService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewProductService(db *sqlx.DB, log l.Logger) *ProductService {
	return &ProductService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}
