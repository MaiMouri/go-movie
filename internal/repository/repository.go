package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
	// GetUserByEmail(email string) (*models.User, error)
	// GetUserByID(id string) (*models.User, error)

	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)
	OneMovie(id int) (*models.Movie, error)
}
