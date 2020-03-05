package password

import (
	"time"

	"github.com/armando0194/movie-night-backend/pkg/api/password"
	"github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// New creates new password logging service
func New(svc password.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents password logging service
type LogService struct {
	password.Service
	logger model.Logger
}

const name = "password"

// Change logging
func (ls *LogService) Change(c *gin.Context, id int, oldPass, newPass string) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Change password request", err,
			[]zap.Field{
				zap.Int("id", id),
				zap.Duration("took", time.Since(begin)),
			},
		)
	}(time.Now())
	return ls.Service.Change(c, id, oldPass, newPass)
}
