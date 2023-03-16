package moviepg

import (
	"log"
	"sesi-8/entity"
	"sesi-8/pkg/errs"
	movierepository "sesi-8/repository/movieRepository"

	"gorm.io/gorm"
)

type moviePG struct {
	db *gorm.DB
}

func NewMoviePG(db *gorm.DB) movierepository.MovieRepository {
	return &moviePG{db: db}
}

func (m *moviePG) CreateMovie(movie entity.Movie) (*entity.Movie, errs.MessageErr) {
	if err := m.db.Create(&movie).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to create movie")
	}

	return &movie, nil
}
