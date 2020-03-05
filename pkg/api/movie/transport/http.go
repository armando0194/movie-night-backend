package transport

import (
	"net/http"
	"strconv"

	"github.com/armando0194/movie-night-backend/pkg/api/movie"
	apperr "github.com/armando0194/movie-night-backend/pkg/utl/error"
	model "github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/gin-gonic/gin"
)

// HTTP represents user http service
type HTTP struct {
	svc movie.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc movie.Service, er *gin.RouterGroup) {
	h := HTTP{svc}
	ur := er.Group("/movie")
	// swagger:route POST /v1/movie movie movieCreate
	// Creates new movie suggestion.
	// responses:
	//  200: movieResp
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
	// ur.GET("/:id", h.view)

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
	// ur.DELETE("/:id", h.delete)
}

// Custom errors
var (
	ErrPasswordsNotMaching = apperr.New(http.StatusBadRequest, "passwords do not match")
)

// User create request
// swagger:model userCreate
type createReq struct {
	Title      string          `json:"Title" validate:"required"`
	Year       string          `json:"Year"`
	Rated      string          `json:"Rated"`
	Released   string          `json:"Released"`
	Runtime    string          `json:"Runtime"`
	Genre      string          `json:"Genre"`
	Director   string          `json:"Director"`
	Writer     string          `json:"Writer"`
	Actors     string          `json:"Actors"`
	Plot       string          `json:"Plot"`
	Language   string          `json:"Language"`
	Country    string          `json:"Country"`
	Awards     string          `json:"Awards"`
	Poster     string          `json:"Poster "`
	Ratings    []model.Ratings `json:"Ratings"`
	Metascore  string          `json:"Metascore"`
	ImdbRating string          `json:"imdbRating"`
	ImdbVotes  string          `json:"imdbVotes"`
	ImdbID     string          `json:"imdbID"`
	Type       string          `json:"Type"`
	DVD        string          `json:"DVD"`
	BoxOffice  string          `json:"BoxOffice"`
	Production string          `json:"Production"`
	Website    string          `json:"Website"`
	Response   string          `json:"Response"`
}

func (h *HTTP) create(c *gin.Context) {
	r := new(createReq)

	if err := c.Bind(r); err != nil {
		apperr.Response(c, err)
		return
	}

	movie, err := h.svc.Create(c, model.Movie{
		Seen:       false,
		Votes:      0,
		Title:      r.Title,
		Year:       r.Year,
		Rated:      r.Rated,
		Released:   r.Released,
		Runtime:    r.Runtime,
		Genre:      r.Genre,
		Director:   r.Director,
		Writer:     r.Writer,
		Actors:     r.Actors,
		Plot:       r.Plot,
		Language:   r.Language,
		Country:    r.Country,
		Awards:     r.Awards,
		Poster:     r.Poster,
		Ratings:    r.Ratings,
		Metascore:  r.Metascore,
		ImdbRating: r.ImdbRating,
		ImdbVotes:  r.ImdbVotes,
		ImdbID:     r.ImdbID,
		Type:       r.Type,
		DVD:        r.DVD,
		BoxOffice:  r.BoxOffice,
		Production: r.Production,
		Website:    r.Website,
		Response:   r.Response,
	})

	if err != nil {
		apperr.Response(c, err)
		return
	}

	c.JSON(http.StatusOK, movie)
}

type listResponse struct {
	Users []model.Movie `json:"movies"`
	Page  int           `json:"page"`
}

func (h *HTTP) list(c *gin.Context) {
	p := new(model.PaginationReq)
	if err := c.ShouldBindQuery(p); err != nil {
		apperr.Response(c, err)
		return
	}

	seen, err := strconv.ParseBool(c.Param("seen"))
	if err != nil {
		seen = false
	}

	result, err := h.svc.List(c, seen, p.Transform())
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

// // User update request
// // swagger:model userUpdate
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

// func (h *HTTP) delete(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		apperr.Response(c, apperr.BadRequest)
// 		return
// 	}

// 	if err := h.svc.Delete(c, id); err != nil {
// 		apperr.Response(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, "Ok")
// }
