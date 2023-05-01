package rest

import (
	"cloud.google.com/go/firestore"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

//counterfeiter:generate . Handler
type Handler interface {
	HealthCheck() func(ctx *gin.Context)
	AddMeal(client *firestore.Client) func(ctx *gin.Context)
}

type client struct {
	Logger *zerolog.Logger
	Router *gin.Engine
}

func New(logger *zerolog.Logger) (*client, error) {

	if logger == nil {
		return nil, errors.New("logger should not be null")
	}

	var instance = client{
		Logger: logger,
	}

	router := gin.Default()
	router.GET("v1/", instance.HealthCheck())
	router.POST("v1/meal", instance.AddMeal())

	instance.Router = router
	return &instance, nil

}
