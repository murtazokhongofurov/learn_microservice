package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	pb "github.com/three_serviceSimple/store_service/genproto/store"

	"github.com/jmoiron/sqlx"
)

type StoreRepo struct {
	db *sqlx.DB
}

func NewStoreRepo(db *sqlx.DB) *StoreRepo {
	return &StoreRepo{
		db: db,
	}
}

func (r *StoreRepo) CreateStore(store *pb.StoreRequest) (*pb.StoreResponse, error) {
	storeResp := pb.StoreResponse{}
	store_id := uuid.New().String()
	err := r.db.QueryRow(`INSERT INTO stores (id,name) values($1,$2) returning id, name`, store_id, store.Name).Scan(&storeResp.Id, &storeResp.Name)
	if err != nil {
		return &pb.StoreResponse{}, err
	}
	var addresses []*pb.AddressResponse
	address_id := uuid.New().String()
	for _, address := range store.Addresses {
		add := pb.AddressResponse{}
		err := r.db.QueryRow(`INSERT INTO 
		addresses 
		(id, store_id, country, street) 
		VALUES($1,$2,$3,$4) 
		RETURNING id, store_id,country, street`,
			address_id, store_id, address.Country, address.Street).Scan(
			&add.Id, &add.StoreId, &add.Country, &add.Street,
		)
		if err != nil {
			return &pb.StoreResponse{}, err
		}
		addresses = append(addresses, &add)
	}
	storeResp.Addresses = addresses
	var products []*pb.ProductRes
	for _, p := range store.Products {
		prod := pb.ProductRes{}
		err := r.db.QueryRow(`INSERT INTO products(
			id, store_id, name, 
			model, price, amount) 
			VALUES($1,$2,$3,$4,$5) 
			RETURNING 
			id,	store_id,name, 
			model, price, amount`,
			uuid.New().String(),
			store_id, p.Name, p.Model,
			p.Price, p.Amount).Scan(
			&prod.Id,&prod.StoreId,
			&prod.Name,	&prod.Model,
			&prod.Price, &prod.Amount)
		if err != nil {
			return &pb.StoreResponse{}, err
		}
		products = append(products, &prod)
	}
	storeResp.Products = products
	return &storeResp, nil
}

func (r *StoreRepo) GetStoreInfo(store *pb.StoreId) (*pb.StoreResponse, error) {
	storeResp := pb.StoreResponse{}
	err := r.db.QueryRow(`SELECT 
		id, 
		name FROM stores WHERE id = $1`, store.StoreId).Scan(&storeResp.Id, &storeResp.Name)
	if err != nil {
		return &pb.StoreResponse{}, err
	}
	rows, err := r.db.Query(`SELECT id, store_id, country, street FROM addresses WHERE store_id=$1`, store.StoreId)
	if err == sql.ErrNoRows {
		return &pb.StoreResponse{}, err
	}
	if err != nil {
		return &pb.StoreResponse{}, err
	}
	for rows.Next() {
		add := pb.AddressResponse{}
		err = rows.Scan(&add.Id, &add.StoreId, &add.Country, &add.Street)
		if err != nil {
			return &pb.StoreResponse{}, err
		}
		storeResp.Addresses = append(storeResp.Addresses, &add)
	}
	return &storeResp, nil
}

func (r *StoreRepo) UpdateStore(store *pb.StoreResponse) (*pb.StoreResponse, error) {
	str := pb.StoreResponse{}
	err := r.db.QueryRow(`update 
	stores set 
	name = $1 
	where id = $2
	RETURNING id, name`, store.Name, store.Id).Scan(
		&str.Id, &str.Name,
	)
	if err != nil {
		return &pb.StoreResponse{}, err
	}
	return &str, nil
}

func (r *StoreRepo) DeleteStore(store *pb.StoreId) (*pb.Empty, error) {
	_, err := r.db.Exec(`delete from stores where id = $1`, store.StoreId)
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}
