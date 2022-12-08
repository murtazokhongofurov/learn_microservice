package storage
import (
	"github.com/three_serviceSimple/product_service/storage/postgres"
	"github.com/three_serviceSimple/product_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Product() repo.ProductStorageI
}

type storagePg struct {
	db *sqlx.DB
	productResp repo.ProductStorageI
}


func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db: db,
		productResp: postgres.NewProductRepo(db),
	}
}

func (s storagePg) Product() repo.ProductStorageI {
	return s.productResp
}
