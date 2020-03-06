package night

import (
	model "github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/gin-gonic/gin"
)

// Create creates a new movie
func (u *MovieNight) Create(c *gin.Context, req model.MovieNight) (*model.MovieNight, error) {
	return u.mndb.Create(u.db, req)
}

// List returns list of movies
func (u *MovieNight) List(c *gin.Context, p *model.Pagination) ([]model.MovieNight, error) {
	return u.mndb.List(u.db, p)
}

func (u *MovieNight) Delete(c *gin.Context, id int) error {
	movie, err := u.mndb.View(u.db, id)
	if err != nil {
		return err
	}

	return u.mndb.Delete(u.db, movie)
}

func (u *MovieNight) AddHost(c *gin.Context, night_id int) (*model.MovieNight, error) {
	movie, err := u.mndb.View(u.db, night_id)
	if err != nil {
		return nil, err
	}

	au := u.rbac.User(c)
	user, err := u.udb.View(u.db, au.ID)

	if err != nil {
		return nil, err
	}

	movie.UpdateHost(user)

	return u.mndb.Update(u.db, movie)
}

func (u *MovieNight) AddRSVP(c *gin.Context, id int) (*model.MovieNight, error) {
	movie, err := u.mndb.View(u.db, id)
	if err != nil {
		return nil, err
	}

	au := u.rbac.User(c)
	user, err := u.udb.View(u.db, au.ID)

	if err != nil {
		return nil, err
	}

	movie.UpdateRSVP(user)

	return u.mndb.Update(u.db, movie)
}
