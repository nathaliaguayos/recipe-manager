package rest_test

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/recipe-manager/cmd/recipe-manager/rest"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	logger := zerolog.Ctx(context.Background())

	tests := map[string]struct {
		Given       *zerolog.Logger
		ExpectError bool
	}{
		"Fail when the given logger is empty": {
			Given:       nil,
			ExpectError: true,
		},
		"Succeed when the given logger is valid": {
			Given:       logger,
			ExpectError: false,
		},
	}
	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			router, err := rest.New(test.Given)
			if test.ExpectError {
				require.Error(t, err)
				require.Nil(t, router)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, router)
		})
	}
}

// NOTE: This works using the env var GOOGLE_APPLICATION_CREDENTIALS
func TestAddMeal(t *testing.T) {
	// create a mock gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// create a mock logger
	mockLogger := zerolog.New(os.Stdout)

	// create a mock client instance
	mockClient, _ := rest.New(&mockLogger)

	// call the AddMeal method with the mock context
	mockClient.AddMeal()(c)

	// check if the expected values are returned
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
