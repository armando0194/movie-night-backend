package query

import (
	apperr "github.com/armando0194/movie-night-backend/pkg/utl/error"
	model "github.com/armando0194/movie-night-backend/pkg/utl/model"
)

// List prepares data for list queries
func List(u *model.AuthUser) (*model.ListQuery, error) {
	switch true {
	case u.Role <= model.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	default:
		return nil, apperr.Forbidden
	}
}
