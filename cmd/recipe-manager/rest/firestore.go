package rest

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

const (
	projectId      string = "inisde-nutrition"
)

func (c *client) CreateClient() (*firestore.Client, error) {
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