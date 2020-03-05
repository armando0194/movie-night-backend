package model

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger represents logging interface
type Logger interface {
	// source, msg, error, params
	Log(ctx *gin.Context, source string, msg string, err error, params []zap.Field)
}
