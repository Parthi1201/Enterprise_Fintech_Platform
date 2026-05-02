package data

import (
	"transaction-service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

)

// ProviderSet is data providers.
var DataProviderSet = wire.NewSet(NewData, NewTransactionRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	DB *gorm.DB
}

// NewData .
func NewData(c *conf.Data) (*Data, func(), error) {
	db,err:=gorm.Open(postgres.Open(c.Database.Source),&gorm.Config{})
	if err!=nil{
		log.Info("Error Reaching database",err)
		return nil,nil,err
	}
	log.Info("database connected")
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{DB:db}, cleanup, nil
}
