package movie

import (
	"github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/gin-gonic/gin"
)

// Create creates a new movie
func (u *Movie) Create(c *gin.Context, req model.Movie) (*model.Movie, error) {
	return u.mdb.Create(u.db, req)
}

// List returns list of movies
func (u *Movie) List(c *gin.Context, seen bool, p *model.Pagination) ([]model.Movie, error) {
	return u.mdb.List(u.db, seen, p)
}
