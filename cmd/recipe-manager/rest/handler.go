package rest

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/recipe-manager/cmd/recipe-manager/meals"
	"net/http"
)

const (
	projectId      string = "inisde-nutrition"
	collectionName string = "meals"
)

func (c *client) HealthCheck() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Ok": true})
	}
}

func (c client) AddMeal() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fsClient, err := c.createClient()
		if err != nil {
			c.Logger.Error().Err(err).Msg("error creating firestore client")
			return
		}
		defer fsClient.Close()
		var meal meals.Meal
		if err := ctx.BindJSON(&meal); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Logger.Info().Fields(meal).Msg("Body deserialized successfully")

		ref := fsClient.Collection(collectionName).NewDoc()

		if meal.ID == "" {
			mealID, err := uuid.NewRandom()
			if err != nil {
				c.Logger.Error().Msg(fmt.Sprintf("failed to create meal ID %v", err))
				return
			}
			meal.ID = mealID.String()
		}

		_, err = ref.Set(ctx, map[string]interface{}{
			"ID":          meal.ID,
			"name":        meal.Name,
			"label":       meal.Label,
			"ingredients": meal.Ingredients,
			"preparation": meal.Procedure,
			"image":       meal.Image,
			"kcal":        meal.Kcal,
			"macros":      meal.Macros,
		})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Logger.Fatal().Err(err).Msg("error creating a new meal")
			return
		}
		ctx.JSON(http.StatusCreated, meal)
	}
}
func (c *client) createClient() (*firestore.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectId}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		c.Logger.Fatal().Err(err).Msg(fmt.Sprintf("error initializing firebase app: %v", err))
		return nil, err
	}

	fsClient, err := app.Firestore(ctx)
	if err != nil {
		c.Logger.Fatal().Err(err).Msg(fmt.Sprintf("error creating firestore client: %v", err))
		return nil, err
	}

	return fsClient, nil
}
