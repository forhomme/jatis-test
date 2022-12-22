package main

import (
	"flag"
	"jatis-test/app/core/models"
	"jatis-test/app/core/usecase"
	"jatis-test/app/infrastructure/database"
	"jatis-test/app/interface/api"
	"jatis-test/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
)

func assignRouting(e *echo.Echo) {
	//assign middlewares
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      nil,
		ErrorMessage: "Timeout Service",
		Timeout:      1 * time.Minute,
	}))

	//api handlers
	apiEngine := e.Group(config.Config.RootURL + "/api")

	// db connection
	dbConnection := config.GetInstancePostgresDb()
	dbRepo := database.NewDBConnection(dbConnection)

	// order api
	orderUc := usecase.NewOrdersUC(dbRepo)
	orderApi := api.NewOrderApi(orderUc)
	orders := apiEngine.Group("/orders")
	orders.GET("/:id", orderApi.GetOrderDetails)
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})

	var migrate bool
	var seeding string
	flag.BoolVar(&migrate, "migrate", true, "If migrate true")
	flag.StringVar(&seeding, "seeding", "", "Seed data from csv file")
	flag.Parse()

	if migrate {
		models.Migrate()
	}

	if len(seeding) > 0 {
		models.Seeding(seeding)
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      AllowOriginSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{"*"},
	}))
	assignRouting(e)

	e.Server.Addr = config.Config.Port
	graceful.ListenAndServe(e.Server, 10*time.Second)
	logrus.Infoln("Start server on port : ", e.Server.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func AllowOriginSkipper(c echo.Context) bool {
	return false
}
