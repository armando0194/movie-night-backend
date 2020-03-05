package rbac

import (
	apperr "github.com/armando0194/movie-night-backend/pkg/utl/error"
	"github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/gin-gonic/gin"
)

// New creates new RBAC service
func New() *Service {
	return &Service{}
}

// Service is RBAC application service
type Service struct{}

func checkBool(b bool) error {
	if b {
		return nil
	}
	return apperr.Forbidden
}

// User returns user data stored in jwt token
func (s *Service) User(c *gin.Context) *model.AuthUser {
	id := c.GetInt("id")
	user := c.GetString("username")
	email := c.GetString("email")
	role := c.MustGet("role").(model.AccessRole)
	return &model.AuthUser{
		ID:       id,
		Username: user,
		Email:    email,
		Role:     role,
	}
}

// EnforceRole authorizes request by AccessRole
func (s *Service) EnforceRole(c *gin.Context, r model.AccessRole) error {
	return checkBool(!(c.MustGet("role").(model.AccessRole) > r))
}

// EnforceUser checks whether the request to change user data is done by the same user
func (s *Service) EnforceUser(c *gin.Context, ID int) error {
	// TODO: Implement querying db and checking the requested user's company_id/location_id
	// to allow company/location admins to view the user
	if s.isAdmin(c) {
		return nil
	}

	return checkBool(c.GetInt("id") == ID)
}

func (s *Service) isAdmin(c *gin.Context) bool {
	return !(c.MustGet("role").(model.AccessRole) > model.AdminRole)
}

// AccountCreate performs auth check when creating a new account
// Location admin cannot create accounts, needs to be fixed on EnforceLocation function
func (s *Service) AccountCreate(c *gin.Context, roleID model.AccessRole) error {
	return s.IsLowerRole(c, roleID)
}

// IsLowerRole checks whether the requesting user has higher role than the user it wants to change
// Used for account creation/deletion
func (s *Service) IsLowerRole(c *gin.Context, r model.AccessRole) error {
	return checkBool(c.MustGet("role").(model.AccessRole) < r)
}
