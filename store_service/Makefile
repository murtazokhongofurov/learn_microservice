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

migrate_store:
	migrate create -ext sql -dir migrations -seq create_stores_table

# migrate_up:
# 	migrate -source file://migrations -database postgres://murtazokhon:12345@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com:5432/storedata?sslmode=disable up

# migrate_down:
# 	migrate -path migrations/ -database postgres://murtazokhon:12345@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com:5432/storedata down

# migrate_force:
# 	migrate -path migrations/ -database postgres://murtazokhon:12345@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com:5432/storedata force 1



migrate_up:
	migrate -source file://migrations -database postgres://developer:2002@localhost:5432/store_service?sslmode=disable up

migrate_down:
	migrate -path migrations/ -database postgres://developer:2002@localhost:5432/store_service down

migrate_force:
	migrate -path migrations/ -database postgres://developer:2002@localhost:5432/store_service force 1
