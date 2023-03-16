package handler

import (
	"net/http"
	"sesi-8/dto"
	"sesi-8/service"

	"github.com/gin-gonic/gin"
)

type movieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) movieHandler {
	return movieHandler{
		movieService: movieService,
	}
}

// CreateNewMovie godoc
// @Tags movies
// @Description Create new movie data
// @ID create-new-movie
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewMovieRequest true "request body json"
// @Success 201 {object} dto.NewMovieResponse
// @Router /movies [post]
func (m *movieHandler) CreateMovie(ctx *gin.Context) {
	var movieRequestBody dto.NewMovieRequest

	if err := ctx.ShouldBindJSON(&movieRequestBody); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	newMovie, err := m.movieService.CreateMovie(movieRequestBody)
	if err != nil {
		ctx.JSON(err.Status(), err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, newMovie)
}
