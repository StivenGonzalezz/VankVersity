package repository

import (
	"video-service/internal/domain/model"
	"video-service/internal/domain/ports"
	"gorm.io/gorm"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
)

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo() ports.VideoRepository{
	dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}
	
	db.AutoMigrate(&model.Video{})

	return &PostgresRepo{db: db}
}
