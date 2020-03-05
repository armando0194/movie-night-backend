package transport

import (
	"net/http"

	apperr "github.com/armando0194/movie-night-backend/pkg/utl/error"

	"github.com/armando0194/movie-night-backend/pkg/api/auth"
	"github.com/gin-gonic/gin"
)

// HTTP represents auth http service
type HTTP struct {
	svc auth.Service
}

// NewAuth creates new auth http service
func NewHTTP(svc auth.Service, r *gin.Engine, mw gin.HandlerFunc) {
	h := HTTP{svc}
	// swagger:route POST /login auth login
	// Logs in user by username and password.
	// responses:
	//  200: loginResp
	//  400: errMsg
	//  401: errMsg
	// 	403: err
	//  404: errMsg
	//  500: err
	r.POST("/login", h.login)
	// swagger:operation GET /refresh/{token} auth refresh
	// ---
	// summary: Refreshes jwt token.
	// description: Refreshes jwt token by checking at database whether refresh token exists.
	// parameters:
	// - name: token
	//   in: path
	//   description: refresh token
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/refreshResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	r.GET("/refresh/:token", h.refresh)

	// swagger:route GET /me auth meReq
	// Gets user's info from session.
	// responses:
	//  200: userResp
	//  500: err
	r.GET("/me", mw, h.me)
}

type credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *HTTP) login(c *gin.Context) {

	cred := new(credentials)
	if err := c.ShouldBindJSON(cred); err != nil {
		apperr.Response(c, err)
		return
	}
	r, err := h.svc.Authenticate(c, cred.Username, cred.Password)

	if err != nil {
		apperr.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, r)
}

func (h *HTTP) refresh(c *gin.Context) {
	r, err := h.svc.Refresh(c, c.Param("token"))
	if err != nil {
		apperr.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, r)
}

func (h *HTTP) me(c *gin.Context) {
	user, err := h.svc.Me(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, error.Error)
		return
	}
	c.JSON(http.StatusOK, user)
}
