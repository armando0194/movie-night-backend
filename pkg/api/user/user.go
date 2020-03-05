// Package user contains user application services
package user

import (
	model "github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/armando0194/movie-night-backend/pkg/utl/query"
	"github.com/gin-gonic/gin"
)

// Create creates a new user account
func (u *User) Create(c *gin.Context, req model.User) (*model.User, error) {
	if err := u.rbac.AccountCreate(c, req.RoleID); err != nil {
		return nil, err
	}
	req.Password = u.sec.Hash(req.Password)
	return u.udb.Create(u.db, req)
}

// List returns list of users
func (u *User) List(c *gin.Context, p *model.Pagination) ([]model.User, error) {
	au := u.rbac.User(c)
	q, err := query.List(au)
	if err != nil {
		return nil, err
	}
	return u.udb.List(u.db, q, p)
}

// View returns single user
func (u *User) View(c *gin.Context, id int) (*model.User, error) {
	if err := u.rbac.EnforceUser(c, id); err != nil {
		return nil, err
	}
	return u.udb.View(u.db, id)
}

// Delete deletes a user
func (u *User) Delete(c *gin.Context, id int) error {
	user, err := u.udb.View(u.db, id)
	if err != nil {
		return err
	}
	if err := u.rbac.IsLowerRole(c, user.Role.AccessLevel); err != nil {
		return err
	}
	return u.udb.Delete(u.db, user)
}

// Update contains user's information used for updating
type Update struct {
	ID        int
	FirstName string
	LastName  string
	Mobile    string
	Phone     string
	Address   string
}

// Update updates user's contact information
func (u *User) Update(c *gin.Context, r *Update) (*model.User, error) {
	if err := u.rbac.EnforceUser(c, r.ID); err != nil {
		return nil, err
	}

	if err := u.udb.Update(u.db, &model.User{
		Base:      model.Base{ID: r.ID},
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Mobile:    r.Mobile,
		Address:   r.Address,
	}); err != nil {
		return nil, err
	}

	return u.udb.View(u.db, r.ID)
}
