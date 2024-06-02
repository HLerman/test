package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	loggerKey       = "logger"
	requestIDHeader = "X-Request-Id"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		logger := log.With().
			Str("request_id", uuid.String()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Logger()

		ctx := context.WithValue(c.Request.Context(), loggerKey, logger)
		c.Request = c.Request.WithContext(ctx)

		c.Set("logger", logger)
		c.Header(requestIDHeader, uuid.String())
		c.Next()
	}
}

func GetLogger(ctx context.Context) zerolog.Logger {
	return ctx.Value(loggerKey).(zerolog.Logger)
}
