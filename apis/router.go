package apis

import (
	"net/http"

	v1 "github.com/ju4n97/gon/apis/v1"
	"github.com/ju4n97/gon/internal/config"
	"github.com/ju4n97/gon/tools/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRequestID: true,
		LogProtocol:  true,
		LogHost:      true,
		LogRemoteIP:  true,
		LogUserAgent: true,
		LogMethod:    true,
		LogURI:       true,
		LogStatus:    true,
		LogLatency:   true,
		LogError:     true,
		LogValuesFunc: func(c echo.Context, req middleware.RequestLoggerValues) error {
			fields := logger.Fields{
				"requestID": req.RequestID,
				"protocol":  req.Protocol,
				"host":      req.Host,
				"ip":        req.RemoteIP,
				"ua":        req.UserAgent,
				"method":    req.Method,
				"uri":       req.URI,
				"status":    req.Status,
				"latency":   req.Latency,
			}

			if req.Error != nil {
				fields["error"] = req.Error.Error()
				logger.Log.Error().WithFields(fields).Msg("Request failed with error")
			} else {
				logger.Log.Info().WithFields(fields).Msg("Request accepted")
			}

			return nil
		},
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.Global.Server.AllowedOrigins,
		AllowMethods: config.Global.Server.AllowedMethods,
		AllowHeaders: config.Global.Server.AllowedHeaders,
	}))
	// e.Use(middleware.Timeout())
	// e.Use(middleware.RateLimiter(&middleware.RateLimiterMemoryStore{}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, friend")
	})

	api := e.Group("/api")

	v1.NewGroup(api)
	// v2.NewGroup(api) // Uncomment this line if you have a v2 group

	return e
}
