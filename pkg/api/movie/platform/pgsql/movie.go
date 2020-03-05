package pgsql

import (
	"net/http"
	"strings"

	apperr "github.com/armando0194/movie-night-backend/pkg/utl/error"
	model "github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/go-pg/pg"

	"github.com/go-pg/pg/orm"
)

// NewMovie returns a new user database instance
func NewMovie() *Movie {
	return &Movie{}
}

// Movie represents the client for user table
type Movie struct{}

// Custom errors
var (
	ErrAlreadyExists = apperr.New(http.StatusInternalServerError, "Movie was already suggested.")
)

// Create creates a new Movie on database
func (u *Movie) Create(db orm.DB, m model.Movie) (*model.Movie, error) {
	var movie = new(model.Movie)
	err := db.Model(movie).Where("lower(title) = ?",
		strings.ToLower(m.Title)).Select()
	if (err == nil) || (err != nil && err != pg.ErrNoRows) {
		return nil, ErrAlreadyExists
	}

	if err := db.Insert(&m); err != nil {
		return nil, err
	}

	return &m, nil
}

// List returns list of all users retrievable for the current user, depending on role
func (u *Movie) List(db orm.DB, seen bool, p *model.Pagination) ([]model.Movie, error) {
	var movies []model.Movie
	q := db.Model(&movies).Limit(p.Limit).Offset(p.Offset).Where("deleted_at is null")

	if err := q.Select(); err != nil {
		return nil, err
	}
	return movies, nil
}
