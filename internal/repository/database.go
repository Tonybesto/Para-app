package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"paramount_school/internal/model"
	"paramount_school/internal/ports"
)

type Postgres struct {
	DB *gorm.DB
}

//NewDB create/returns a new instance of our Database
func NewDB(DB *gorm.DB) ports.Repository {
	return &Postgres{
		DB: DB,
	}
}

//Initialize opens the database, creates jobs table if not created and populate it if its empty and returns a DB
func Initialize(dbURI string) (*gorm.DB, error) {

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("We can't open a DATABASE: %v\n", err)

	}

	conn.AutoMigrate(&model.HighSchool{}, &model.Student{}, &model.Certificate{}, &model.Transcript{},
		&model.Photo{}, &model.LocalId{}, &model.Passport{}, &model.Blacklist{})

	return conn, nil
}
