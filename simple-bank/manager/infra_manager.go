package manager

import (
	"fmt"

	"github.com/rizkyfazri23/mnc/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	DB() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
	cfg config.AppConfig
}

func (i *infraManager) initDb() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name, i.cfg.SslMode)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	i.db = db
	fmt.Println("DB Connected")
}

func (i *infraManager) DB() *gorm.DB {
	return i.db
}

func NewInfraManager(cfg config.AppConfig) InfraManager {
	infra := infraManager{
		cfg: cfg,
	}
	infra.initDb()
	return &infra
}
