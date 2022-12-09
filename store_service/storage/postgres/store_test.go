package postgres

// import (
// 	"gitlab.com/go_kafka/store_service/config"
// 	pb "gitlab.com/go_kafka/store_service/genproto/store"
// 	"gitlab.com/go_kafka/store_service/pkg/db"
// 	"gitlab.com/go_kafka/store_service/storage/repo"

// 	"github.com/stretchr/testify/suite"
// )

// type StoreSuiteTest struct {
// 	suite.Suite
// 	CleanUpfunc func()
// 	Repository  repo.StoreStorageI
// }

// func (s *StoreSuiteTest) SetubSuite() {
// 	pgPool, cleanUpfunc := db.ConnectToDBForSuite(config.Load())
// 	s.Repository = NewStoreRepo(pgPool)
// 	s.CleanUpfunc = cleanUpfunc
// }

// func (s *StoreSuiteTest) TestStoreCrub() {
// 	storeCreate := pb.StoreRequest{
// 		Name: "testing store",
// 	}
// 	store, err := s.Repository.CreateStore(&storeCreate)
// 	s.Nil(err)
// 	s.NotEmpty(store)

// 	updateStore := pb.StoreResponse{
// 		Id:   store.Id,
// 		Name: "update store",
// 	}
// 	update_store, err := s.Repository.UpdateStore(&updateStore)
// 	s.Nil(err)
// 	s.NotEmpty(update_store)

// 	getStore, err := s.Repository.GetStoreInfo(&pb.StoreId{StoreId: update_store.Id})
// 	s.Nil(err)
// 	s.NotEmpty(getStore)
// 	s.Equal(update_store.Name, getStore.Name)

// 	deleteStore, err := s.Repository.DeleteStore(&pb.StoreId{StoreId: getStore.Id})
// 	s.Nil(err)
// 	s.NotEmpty(deleteStore)

// }
