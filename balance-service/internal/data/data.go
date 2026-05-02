package data

import (
	"balance-service/internal/conf"
	"gorm.io/driver/postgres"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBalanceRepo)

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
	if err:=db.AutoMigrate(&Balance{});err!=nil{
		return nil,nil,err
	}
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{DB: db}, cleanup, nil
}
