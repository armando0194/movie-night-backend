package user

import (
	"github.com/armando0194/movie-night-backend/pkg/api/user/platform/pgsql"
	"github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Service represents user application interface
type Service interface {
	Create(*gin.Context, model.User) (*model.User, error)
	List(*gin.Context, *model.Pagination) ([]model.User, error)
	View(*gin.Context, int) (*model.User, error)
	Delete(*gin.Context, int) error
	Update(*gin.Context, *Update) (*model.User, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *User {
	return &User{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *User {
	return New(db, pgsql.NewUser(), rbac, sec)
}

// User represents user application service
type User struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, model.User) (*model.User, error)
	View(orm.DB, int) (*model.User, error)
	List(orm.DB, *model.ListQuery, *model.Pagination) ([]model.User, error)
	Update(orm.DB, *model.User) error
	Delete(orm.DB, *model.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(*gin.Context) *model.AuthUser
	EnforceUser(*gin.Context, int) error
	AccountCreate(*gin.Context, model.AccessRole) error
	IsLowerRole(*gin.Context, model.AccessRole) error
}
