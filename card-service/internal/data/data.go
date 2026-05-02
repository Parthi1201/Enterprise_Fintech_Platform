package data

import (
	"card-service/internal/conf"
	"gorm.io/driver/postgres"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var DataProviderSet = wire.NewSet(NewData,NewCardRepo)

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
	// if err:=db.AutoMigrate(&mdCard{});err!=nil{
	// 	return nil,nil,err
	// }
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{DB: db}, cleanup, nil
}
