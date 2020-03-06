package movie

import (
	"fmt"

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

// Vote adds one to votes counter
func (u *Movie) Vote(c *gin.Context, id int) error {

	movie, err := u.mdb.View(u.db, id)
	if err != nil {
		return err
	}

	movie.IncrementVote()
	fmt.Printf("%#v", movie)
	return u.mdb.Update(u.db, movie)
}

func (u *Movie) Delete(c *gin.Context, id int) error {
	movie, err := u.mdb.View(u.db, id)
	if err != nil {
		return err
	}

	return u.mdb.Delete(u.db, movie)
}
