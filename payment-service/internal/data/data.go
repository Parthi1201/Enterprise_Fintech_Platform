package data

import (
	"payment-service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewPaymentRepo,
)

// Data 
type Data struct {
	db *gorm.DB
}

// NewData 
func NewData(c *conf.Data) (*Data, func(), error) {
	db, err := gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	log.Info("database connected")

	cleanup := func() {
		log.Info("closing data resources")
	}

	return &Data{db: db}, cleanup, nil
}

func NewDB(d *Data) *gorm.DB {
	return d.db
}
