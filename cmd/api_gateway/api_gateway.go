package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/esvm/if1015covidproject-api/src/api_gateway"
	custom_middleware "github.com/esvm/if1015covidproject-api/src/api_gateway/middleware"
	"github.com/esvm/if1015covidproject-api/src/covid_report_service"
	"github.com/esvm/if1015covidproject-api/src/logger_builder"
	"github.com/esvm/if1015covidproject-api/src/metrics_instrumenter"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rollbar/rollbar-go"
)

func setupAppAndMiddlewares(logger log.Logger) *echo.Echo {
	app := echo.New()
	app.Use(custom_middleware.LoggerWithConfig(
		custom_middleware.LoggerConfig{
			Logger: logger,
		},
	))

	app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	maxAge, err := strconv.Atoi(os.Getenv("API_GATEWAY_CORS_MAX_AGE"))
	if err != nil {
		level.Error(logger).Log("err", err, "message", "API_GATEWAY_CORS_MAX_AGE should be an integer")
	}
	origins := strings.Split(os.Getenv("API_GATEWAY_ALLOWED_ORIGINS"), ",")

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: origins,
		AllowMethods: []string{echo.OPTIONS, echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
		MaxAge:       maxAge,
	}))

	app.Use(custom_middleware.ContextInjector())

	app.Use(custom_middleware.RequestID())

	return app
}

func main() {
	logger := logger_builder.NewLogger("api-gateway")

	metrics_instrumenter.Register()

	app := setupAppAndMiddlewares(logger)

	clients := api_gateway.Clients{}
	clients.CovidReportService = covid_report_service.NewCovidReportService(logger)

	api_gateway.MakeRoutes(app, clients)

	port := os.Getenv("API_GATEWAY_PORT")
	if port == "" {
		port = "80"
	}

	go func() {
		app.Logger.Info("starting the server")

		if err := app.Start(":" + port); err != nil {
			app.Logger.Errorf("shutting down the server: %s", err)
		}
	}()

	go func() {
		rollbar.Wait()
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds, any request that lasts more than that
	// will be dropped, looking as 504 (GATEWAY TIMEOUT) to the client
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}
}
