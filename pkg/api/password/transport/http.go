package transport

import (
	"net/http"
	"strconv"

	apperr "github.com/armando0194/movie-night-backend/pkg/utl/error"

	"github.com/armando0194/movie-night-backend/pkg/api/password"
	"github.com/gin-gonic/gin"
)

// HTTP represents password http transport service
type HTTP struct {
	svc password.Service
}

// NewHTTP create new password http service
func NewHTTP(svc password.Service, r *gin.RouterGroup) {
	h := HTTP{svc}
	pr := r.Group("/password")

	// swagger:operation PATCH /v1/password/{id} password pwChange
	// ---
	// summary: Changes user's password.
	// description: If user's old passowrd is correct, it will be replaced with new password.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/pwChange"
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	pr.PATCH("/:id", h.change)
}

// Custom errors
var (
	ErrPasswordsNotMaching = apperr.New(http.StatusBadRequest, "passwords do not match")
)

// Password change request
// swagger:model pwChange
type changeReq struct {
	ID                 int    `json:"-"`
	OldPassword        string `json:"old_password" validate:"required,min=8"`
	NewPassword        string `json:"new_password" validate:"required,min=8"`
	NewPasswordConfirm string `json:"new_password_confirm" validate:"required"`
}

func (h *HTTP) change(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apperr.Response(c, err)
		return
	}

	p := new(changeReq)
	if err := c.Bind(p); err != nil {
		apperr.Response(c, err)
		return
	}

	if p.NewPassword != p.NewPasswordConfirm {
		apperr.Response(c, err)
		return
	}

	if err := h.svc.Change(c, id, p.OldPassword, p.NewPassword); err != nil {
		apperr.Response(c, err)
		return
	}

	c.JSON(http.StatusOK, "OK")
}
