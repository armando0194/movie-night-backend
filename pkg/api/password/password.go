package password

import (
	"net/http"

	apperr "github.com/armando0194/movie-night-backend/pkg/utl/error"
	"github.com/gin-gonic/gin"
)

// Custom Errors
var (
	ErrorIncorrectPassword = apperr.New(http.StatusBadRequest, "Incorrect old password")
	ErrorInsecurePassword  = apperr.New(http.StatusBadRequest, "Insecure password")
)

// Change changes user's password
func (p *Password) Change(c *gin.Context, userID int, oldPass, newPass string) error {
	if err := p.rbac.EnforceUser(c, userID); err != nil {
		return err
	}

	u, err := p.udb.View(p.db, userID)
	if err != nil {
		return err
	}

	if !p.sec.HashMatchesPassword(u.Password, oldPass) {
		return ErrorIncorrectPassword
	}

	if !p.sec.Password(newPass, u.FirstName, u.LastName, u.Username, u.Email) {
		return ErrorInsecurePassword
	}

	u.ChangePassword(p.sec.Hash(newPass))

	return p.udb.Update(p.db, u)
}
