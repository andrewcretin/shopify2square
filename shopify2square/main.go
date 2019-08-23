package main

import (
	"context"
	"github.com/andrewcretin/shopify2square/controller"
	"github.com/andrewcretin/shopify2square/models"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(_ context.Context) (*models.SyncResponse, error) {

	c, err := controller.NewController()
	if err != nil {
		return nil, err
	}

	// Perform action
	response, err := c.SyncShopifyToSquare()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
