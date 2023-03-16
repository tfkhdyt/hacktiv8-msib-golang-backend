package movierepository

import (
	"sesi-8/entity"
	"sesi-8/pkg/errs"
)

type MovieRepository interface {
	CreateMovie(entity.Movie) (*entity.Movie, errs.MessageErr)
}
