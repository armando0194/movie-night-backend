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
	q := db.Model(&movies).Limit(p.Limit).Offset(p.Offset).Where("seen = ?", seen)

	if err := q.Select(); err != nil {
		return nil, err
	}
	return movies, nil
}

// View returns single movie by ID
func (u *Movie) View(db orm.DB, id int) (*model.Movie, error) {
	var movie = new(model.Movie)
	sql := `SELECT * FROM "movies" WHERE ("movies"."id" = ? and deleted_at is null)`
	_, err := db.QueryOne(movie, sql, id)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

// Update updates user's info
func (u *Movie) Update(db orm.DB, movie *model.Movie) error {
	return db.Update(movie)
}

// Delete sets deleted_at for a user
func (u *Movie) Delete(db orm.DB, movie *model.Movie) error {
	return db.Delete(movie)
}
