package persistence

import (
	"coupon-service/config"
	"coupon-service/domains/entities"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbConn(dbConfig *config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Port, dbConfig.DatabaseName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to postgres: %s", err)
	}

	if dbConfig.IsMigration {
		migrateDB(db)
	}

	return db
}

func migrateDB(db *gorm.DB) {
	fmt.Println("Migrating database...")
	db.AutoMigrate(
		&entities.BoardUser{},
		&entities.Board{},
		&entities.BoardMember{},
		&entities.Coupon{},
		&entities.CouponUsageHistory{},
	)
	fmt.Println("Migrating Success!")
}
