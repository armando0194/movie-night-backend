package user

import (
	"fmt"
	"time"

	"github.com/armando0194/movie-night-backend/pkg/api/movie"
	"github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// New creates new user logging service
func New(svc movie.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents user logging service
type LogService struct {
	movie.Service
	logger model.Logger
}

const name = "user"

// Create logging
func (ls *LogService) Create(c *gin.Context, req model.Movie) (resp *model.Movie, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Create user request", err,
			[]zap.Field{
				zap.String("req", fmt.Sprintf("%#v", req)),
				zap.String("resp", fmt.Sprintf("%#v", resp)),
				zap.Duration("took", time.Since(begin)),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, req)
}

// // List logging
// func (ls *LogService) List(c *gin.Context, req *model.Pagination) (resp []model.User, err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "List user request", err,
// 			[]zap.Field{
// 				zap.String("req", fmt.Sprintf("%#v", req)),
// 				zap.String("resp", fmt.Sprintf("%#v", resp)),
// 				zap.Duration("took", time.Since(begin)),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.List(c, req)
// }

// // View logging
// func (ls *LogService) View(c *gin.Context, req int) (resp *model.User, err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "View user request", err,
// 			[]zap.Field{
// 				zap.Int("req", req),
// 				zap.String("resp", fmt.Sprintf("%#v", resp)),
// 				zap.Duration("took", time.Since(begin)),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.View(c, req)
// }

// // Delete logging
// func (ls *LogService) Delete(c *gin.Context, req int) (err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "Delete user request", err,
// 			[]zap.Field{
// 				zap.Int("req", req),
// 				zap.Duration("took", time.Since(begin)),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.Delete(c, req)
// }

// // Update logging
// func (ls *LogService) Update(c *gin.Context, req *user.Update) (resp *model.User, err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "Update user request", err,
// 			[]zap.Field{
// 				zap.String("req", fmt.Sprintf("%#v", req)),
// 				zap.String("resp", fmt.Sprintf("%#v", resp)),
// 				zap.Duration("took", time.Since(begin)),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.Update(c, req)
// }
