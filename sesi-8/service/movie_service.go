package service

import (
	"net/http"
	"sesi-8/dto"
	"sesi-8/entity"
	"sesi-8/pkg/errs"
	movierepository "sesi-8/repository/movieRepository"
)

type MovieService interface {
	CreateMovie(payload dto.NewMovieRequest) (*dto.NewMovieResponse, errs.MessageErr)
}

type movieService struct {
	movieRepo movierepository.MovieRepository
}

func NewMovieService(movieRepo movierepository.MovieRepository) MovieService {
	return &movieService{movieRepo: movieRepo}
}

func (m *movieService) CreateMovie(
	payload dto.NewMovieRequest,
) (*dto.NewMovieResponse, errs.MessageErr) {
	movie := entity.Movie{
		Title: payload.Title,
		Price: payload.Price,
	}

	newMovie, err := m.movieRepo.CreateMovie(movie)
	if err != nil {
		return nil, err
	}

	response := &dto.NewMovieResponse{
		Message:    "success",
		StatusCode: http.StatusCreated,
		Data: dto.MovieResponse{
			Id:    newMovie.ID,
			Title: newMovie.Title,
			Price: newMovie.Price,
		},
	}

	return response, nil
}
