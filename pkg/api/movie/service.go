package movie

import (
	"github.com/armando0194/movie-night-backend/pkg/api/movie/platform/pgsql"
	model "github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Service represents Movie application interface
type Service interface {
	Create(*gin.Context, model.Movie) (*model.Movie, error)
	List(*gin.Context, bool, *model.Pagination) ([]model.Movie, error)
	// View(*gin.Context, int) (*model.Movie, error)
	Delete(*gin.Context, int) error
	Vote(*gin.Context, int) error
}

// New creates new Movie application service
func New(db *pg.DB, mdb MDB) *Movie {
	return &Movie{db: db, mdb: mdb}
}

// Initialize initalizes Movie application service with defaults
func Initialize(db *pg.DB) *Movie {
	return New(db, pgsql.NewMovie())
}

// Movie represents movie application service
type Movie struct {
	db  *pg.DB
	mdb MDB
}

// MDB represents Movie repository interface
type MDB interface {
	Create(orm.DB, model.Movie) (*model.Movie, error)
	View(orm.DB, int) (*model.Movie, error)
	List(orm.DB, bool, *model.Pagination) ([]model.Movie, error)
	// Vote(orm.DB, *model.Movie) error
	Update(orm.DB, *model.Movie) error
	Delete(orm.DB, *model.Movie) error
}
