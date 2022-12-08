package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	pb "github.com/three_serviceSimple/product_service/genproto/product"
	pbp "github.com/three_serviceSimple/product_service/genproto/store"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *productRepo {
	return &productRepo{db: db}
}

func (r *productRepo) CreateProduct(product *pbp.ProductRes) error {
	fmt.Println("created product")
	_, err := r.db.Exec(`INSERT INTO 
	products(
		id, 
		store_id, 
		name, model, 
		price, amount) VALUES($1,$2,$3,$4,$5,$6)`,
		product.Id,
		product.StoreId,
		product.Name,
		product.Model,
		product.Price,
		product.Amount)
	if err != nil {
		fmt.Println("error while consumer products from store")
		return err
	}
	return nil
}

func (r *productRepo) GetProductById(product *pb.ProductId) (*pb.Product, error) {
	response := pb.Product{}
	err := r.db.QueryRow(`SELECT 
	id, 
	store_id,
	name, 
	model, 
	price, 
	amount 
	FROM products WHERE id = $1`, product.Id).Scan(
		&response.Id,
		&response.StoreId,
		&response.Name,
		&response.Model,
		&response.Price,
		&response.Amount,
	)
	if err != nil {
		return &pb.Product{}, err
	}
	return &response, nil
}

func (r *productRepo) UpdateProduct(product *pb.Product) (*pb.Product, error) {
	response := pb.Product{}
	_, err := r.db.Exec(`UPDATE products SET 
	name = $1, 
	model = $2,
	price=$3,
	amount=$4
	WHERE id = $5`, product.Name, product.Model, product.Price, product.Amount, product.Id)
	if err != nil {
		return &pb.Product{}, err
	}
	return &response, nil
}

func (r *productRepo) DeleteProduct(product *pb.ProductId) (*pb.Empty, error) {
	_, err := r.db.Exec(`delete from products where id = $1`, product.Id)
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}
