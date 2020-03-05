package zap

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Log represents zerolog logger
type Log struct {
	logger *zap.Logger
}

// New instantiates new zero logger
func New() *Log {
	z, _ := zap.NewDevelopment()
	return &Log{
		logger: z,
	}
}

// Log logs using zap
func (z *Log) Log(ctx *gin.Context, source string, msg string, err error, params []zap.Field) {

	if params == nil {
		// params = make(map[string]interface{})
		params = make([]zap.Field, 0, 4)
	}

	params = append(params, zap.String("Source", source))

	if _, ok := ctx.Get("id"); ok {
		params = append(params, zap.Int("id", ctx.GetInt("id")))
	}

	if _, ok := ctx.Get("username"); ok {
		params = append(params, zap.String("username", ctx.GetString("username")))
	}

	if err != nil {
		params = append(params, zap.String("error", err.Error()))
		z.logger.Error(msg, params...)
		return
	}

	z.logger.Info(msg, params...)
}
