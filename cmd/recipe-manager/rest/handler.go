package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/recipe-manager/cmd/recipe-manager/meals"
)

const (
	collectionName string = "meals"
)

//HealthCheck review the service status
func (c *client) HealthCheck() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Ok": true})
	}
}

//AddMeal add a new meal into the firestore database
func (c client) AddMeal() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fsClient, err := c.CreateClient()
		if err != nil {
			c.Logger.Error().Err(err).Msg("error creating firestore client")
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer fsClient.Close()
		var meal meals.Meal
		if err := ctx.BindJSON(&meal); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Logger.Info().Fields(meal).Msg("Body deserialized successfully")

		ref := fsClient.Collection(collectionName).NewDoc()

		if meal.ID == "" {
			mealID, err := uuid.NewRandom()
			if err != nil {
				c.Logger.Error().Msg(fmt.Sprintf("failed to create meal ID %v", err))
				ctx.AbortWithError(http.StatusInternalServerError, err)
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

//GetMeals retrieves all meals stored in the firestore database
func (c client) GetMeals() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fsClient, err := c.CreateClient()
		if err != nil {
			c.Logger.Error().Err(err).Msg("error creating firestore client")
			return
		}
		defer fsClient.Close()
		
		mealsCollection := fsClient.Collection(collectionName)
		snapshot, err := mealsCollection.Documents(context.Background()).GetAll()
		if err != nil {
			c.Logger.Error().Err(err).Msg("error creating collection snapshot")
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var allMeals []meals.Meal
		for _, doc := range snapshot {
			var user meals.Meal
			if err := doc.DataTo(&user); err != nil {
				c.Logger.Error().Err(err).Msg("error reading data")
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			allMeals = append(allMeals, user)
		}

		ctx.JSON(http.StatusOK, allMeals)
	}
}
