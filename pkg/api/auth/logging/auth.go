package auth

import (
	"fmt"
	"time"

	"github.com/armando0194/movie-night-backend/pkg/api/auth"
	"github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// New creates new auth logging service
func New(svc auth.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents auth logging service
type LogService struct {
	auth.Service
	logger model.Logger
}

const name = "auth"

// Authenticate logging
func (ls *LogService) Authenticate(c *gin.Context, user string, password string) (resp *model.AuthToken, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Authenticate request", err,
			[]zap.Field{
				zap.String("req", user),
				zap.Duration("took", time.Since(begin)),
			},
		)
	}(time.Now())
	return ls.Service.Authenticate(c, user, password)
}

// Refresh logging
func (ls *LogService) Refresh(c *gin.Context, req string) (resp *model.RefreshToken, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Refresh request", err,
			[]zap.Field{
				zap.String("req", req),
				zap.String("resp", resp.Token),
				zap.Duration("took", time.Since(begin)),
			},
		)
	}(time.Now())
	return ls.Service.Refresh(c, req)
}

// Me logging
func (ls *LogService) Me(c *gin.Context) (resp *model.User, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Me request", err,
			[]zap.Field{
				zap.String("resp", fmt.Sprintf("%#v", resp)),
				zap.Duration("took", time.Since(begin)),
			},
		)
	}(time.Now())
	return ls.Service.Me(c)
}
