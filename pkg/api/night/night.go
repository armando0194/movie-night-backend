package night

import (
	"fmt"

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

func (u *MovieNight) AddHost(c *gin.Context, night_id int) error {
	movie, err := u.mndb.View(u.db, night_id)
	if err != nil {
		return err
	}
	fmt.Printf("%#v", movie)
	au := u.rbac.User(c)
	fmt.Printf("%#v", au)
	user, err := u.udb.View(u.db, au.ID)
	fmt.Printf("%#v", user)
	if err != nil {
		return err
	}

	movie.UpdateHost(user)

	return u.mndb.Update(u.db, movie)
}

func (u *MovieNight) RSVP(c *gin.Context, id int) error {
	movie, err := u.mndb.View(u.db, id)
	if err != nil {
		return err
	}

	return u.mndb.Delete(u.db, movie)
}
