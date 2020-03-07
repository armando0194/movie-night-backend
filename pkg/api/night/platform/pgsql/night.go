package pgsql

import (
	"net/http"

	apperr "github.com/armando0194/movie-night-backend/pkg/utl/error"
	model "github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/go-pg/pg"

	"github.com/go-pg/pg/orm"
)

// NewMovieNight returns a new user database instance
func NewMovieNight() *MovieNight {
	return &MovieNight{}
}

// Movie represents the client for user table
type MovieNight struct{}

// Custom errors
var (
	ErrAlreadyExists = apperr.New(http.StatusInternalServerError, "MovieNight has the same week number.")
)

// Create creates a new Movie on database
func (u *MovieNight) Create(db orm.DB, m model.MovieNight) (*model.MovieNight, error) {
	var movie = new(model.MovieNight)
	err := db.Model(movie).Where("week_number = ?",
		m.WeekNumber).Select()
	if (err == nil) || (err != nil && err != pg.ErrNoRows) {
		return nil, ErrAlreadyExists
	}

	if err := db.Insert(&m); err != nil {
		return nil, err
	}

	return &m, nil
}

// List returns list of all users retrievable for the current user, depending on role
func (u *MovieNight) List(db orm.DB, p *model.Pagination) ([]model.MovieNight, error) {
	var nights []model.MovieNight
	q := db.Model(&nights).Limit(p.Limit).Offset(p.Offset)

	if err := q.Select(); err != nil {
		return nil, err
	}
	return nights, nil
}

// View returns single movie by ID
func (u *MovieNight) View(db orm.DB, id int) (*model.MovieNight, error) {
	var night = new(model.MovieNight)
	sql := `SELECT * FROM "movie_nights" WHERE ("movie_nights"."id" = ? and deleted_at is null)`
	_, err := db.QueryOne(night, sql, id)
	if err != nil {
		return nil, err
	}

	return night, nil
}

// Update updates user's info
func (u *MovieNight) Update(db orm.DB, night *model.MovieNight) (*model.MovieNight, error) {
	return night, db.Update(night)
}

// Delete sets deleted_at for a user
func (u *MovieNight) Delete(db orm.DB, night *model.MovieNight) error {
	return db.Delete(night)
}
