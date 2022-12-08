pull_submodule:
	git submodule update --init --recursive

update_submodule:	
	git submodule update --remote --merge
	
run:
	go run cmd/main.go

proto-gen:
	./script/gen-proto.sh

swag:
	swag init -g ./api/router.go -o api/docs

migrate_product:
	migrate create -ext sql -dir migrations -seq create_products_table

migrate_up:
	migrate -source file://migrations -database postgres://murtazokhon:12345@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com:5432/productdata?sslmode=disable up

migrate_down:
	migrate -path migrations/ -database postgres://murtazokhon:12345@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com:5432/productdata down

migrate_force:
	migrate -path migrations/ -database postgres://murtazokhon:12345@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com:5432/productdata force 1
