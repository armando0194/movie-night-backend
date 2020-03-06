package night

import (
	"github.com/armando0194/movie-night-backend/pkg/api/night/platform/pgsql"
	user "github.com/armando0194/movie-night-backend/pkg/api/user"
	user_psql "github.com/armando0194/movie-night-backend/pkg/api/user/platform/pgsql"
	model "github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Service represents Movie application interface
type Service interface {
	Create(*gin.Context, model.MovieNight) (*model.MovieNight, error)
	List(*gin.Context, *model.Pagination) ([]model.MovieNight, error)
	Delete(*gin.Context, int) error
	AddHost(*gin.Context, int) error
}

// New creates new Movie application service
func New(db *pg.DB, mndb MNDB, udb user.UDB) *MovieNight {
	return &MovieNight{db: db, mndb: mndb}
}

// Initialize initalizes Movie application service with defaults
func Initialize(db *pg.DB) *MovieNight {
	return New(db, pgsql.NewMovieNight(), user_psql.NewUser())
}

// Movie represents movie application service
type MovieNight struct {
	db   *pg.DB
	mndb MNDB
	udb  user.UDB
	rbac RBAC
}

// MDB represents Movie repository interface
type MNDB interface {
	Create(orm.DB, model.MovieNight) (*model.MovieNight, error)
	View(orm.DB, int) (*model.MovieNight, error)
	List(orm.DB, *model.Pagination) ([]model.MovieNight, error)
	Update(orm.DB, *model.MovieNight) error
	Delete(orm.DB, *model.MovieNight) error
}

type RBAC interface {
	User(*gin.Context) *model.AuthUser
}
