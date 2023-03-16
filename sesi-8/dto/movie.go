package dto

type NewMovieRequest struct {
	Title string `json:"title"`
	Price uint   `json:"price"`
}

type MovieResponse struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Price uint   `json:"price"`
}

type NewMovieResponse struct {
	Message    string        `json:"message"`
	StatusCode uint          `json:"statusCode"`
	Data       MovieResponse `json:"data"`
}
