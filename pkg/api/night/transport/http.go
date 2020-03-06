package transport

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/armando0194/movie-night-backend/pkg/api/night"
	apperr "github.com/armando0194/movie-night-backend/pkg/utl/error"
	model "github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/gin-gonic/gin"
)

// HTTP represents user http service
type HTTP struct {
	svc night.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc night.Service, er *gin.RouterGroup) {
	h := HTTP{svc}
	ur := er.Group("/night")
	// swagger:route POST /v1/users users userCreate
	// Creates new user account.
	// responses:
	//  200: userResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	ur.POST("", h.create)

	// swagger:operation GET /v1/users users listUsers
	// ---
	// summary: Returns list of users.
	// description: Returns list of users. Depending on the user role requesting it, it may return all users for SuperAdmin/Admin users, all company/location users for Company/Location admins, and an error for non-admin users.
	// parameters:
	// - name: limit
	//   in: query
	//   description: number of results
	//   type: int
	//   required: false
	// - name: page
	//   in: query
	//   description: page number
	//   type: int
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/userListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("", h.list)

	// swagger:operation GET /v1/users/{id} users getUser
	// ---
	// summary: Returns a single user.
	// description: Returns a single user by its ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/userResp"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "404":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.PUT("/host/:night_id/", h.host)

	// swagger:operation PATCH /v1/users/{id} users userUpdate
	// ---
	// summary: Updates user's contact information
	// description: Updates user's contact information -> first name, last name, mobile, phone, address.
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
	//     "$ref": "#/definitions/userUpdate"
	// responses:
	//   "200":
	//     "$ref": "#/responses/userResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	// ur.PATCH("/:id", h.update)

	// swagger:operation DELETE /v1/users/{id} users userDelete
	// ---
	// summary: Deletes a user
	// description: Deletes a user with requested ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.DELETE("/:id", h.delete)
}

// Custom errors
var (
	ErrPasswordsNotMaching = apperr.New(http.StatusBadRequest, "passwords do not match")
)

// User create request
// swagger:model userCreate
type createReq struct {
	WeekNumber int    `json:"week_number" validate:"required"`
	Date       string `json:"date" validate:"required"`
}

func (h *HTTP) create(c *gin.Context) {
	r := new(createReq)

	if err := c.Bind(r); err != nil {
		apperr.Response(c, err)
		return
	}

	usr, err := h.svc.Create(c, model.MovieNight{
		WeekNumber: r.WeekNumber,
		Date:       r.Date,
	})

	if err != nil {
		apperr.Response(c, err)
		return
	}

	c.JSON(http.StatusOK, usr)
}

type listResponse struct {
	Users []model.MovieNight `json:"nights"`
	Page  int                `json:"page"`
}

func (h *HTTP) list(c *gin.Context) {
	p := new(model.PaginationReq)
	if err := c.ShouldBindQuery(p); err != nil {
		apperr.Response(c, err)
		return
	}

	result, err := h.svc.List(c, p.Transform())

	if err != nil {
		apperr.Response(c, err)
		return
	}

	c.JSON(http.StatusOK, listResponse{result, p.Page})
}

// func (h *HTTP) view(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		apperr.Response(c, apperr.BadRequest)
// 		return

// 	}

// 	result, err := h.svc.View(c, id)
// 	if err != nil {
// 		apperr.Response(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, result)
// }

// User update request
// swagger:model userUpdate
// type updateReq struct {
// 	ID        int    `json:"-"`
// 	FirstName string `json:"first_name,omitempty" validate:"omitempty,min=2"`
// 	LastName  string `json:"last_name,omitempty" validate:"omitempty,min=2"`
// 	Mobile    string `json:"mobile,omitempty"`
// 	Phone     string `json:"phone,omitempty"`
// 	Address   string `json:"address,omitempty"`
// }

// func (h *HTTP) update(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		apperr.Response(c, apperr.BadRequest)
// 		return
// 	}

// 	req := new(updateReq)
// 	if err := c.Bind(req); err != nil {
// 		apperr.Response(c, err)
// 		return
// 	}

// 	usr, err := h.svc.Update(c, &user.Update{
// 		ID:        id,
// 		FirstName: req.FirstName,
// 		LastName:  req.LastName,
// 		Mobile:    req.Mobile,
// 		Phone:     req.Phone,
// 		Address:   req.Address,
// 	})

// 	if err != nil {
// 		apperr.Response(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, usr)
// }

func (h *HTTP) host(c *gin.Context) {
	night_id, err := strconv.Atoi(c.Param("night_id"))
	if err != nil {
		apperr.Response(c, apperr.BadRequest)
		return
	}

	fmt.Println(night_id)
	err = h.svc.AddHost(c, night_id)

	if err != nil {
		apperr.Response(c, err)
		return
	}

	c.JSON(http.StatusOK, night_id)
}

func (h *HTTP) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apperr.Response(c, apperr.BadRequest)
		return
	}

	if err := h.svc.Delete(c, id); err != nil {
		apperr.Response(c, err)
		return
	}

	c.JSON(http.StatusOK, "Ok")
}
