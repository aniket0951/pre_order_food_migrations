package connections

import (
	"fmt"
	"log"
	"pre_order_migration/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func Connect() {
	if PostgresDB == nil {
		var err error
		dsn := "host=" + utils.Config().DBConfig.Host +
			" user=" + utils.Config().DBConfig.Username +
			" password=" + utils.Config().DBConfig.Password +
			" dbname=" + utils.Config().DBConfig.Database +
			" port=" + utils.Config().DBConfig.Port +
			" sslmode=disable"

		PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("[Connection], Error in opening db")
		}
		fmt.Println("DB connection established")
	}
}

func DB() *gorm.DB {
	if PostgresDB == nil {
		Connect()
	}
	return PostgresDB
}
