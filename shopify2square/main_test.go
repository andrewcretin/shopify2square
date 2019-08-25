package main

import (
	"context"
	"fmt"
	"github.com/andrewcretin/shopify2square/test"
	"testing"
)

func TestHandler_test(t *testing.T) {

	err := test.SetupTestEnvironment(t)
	if err != nil {
		t.Errorf("\n setupTestEnvironment error: %+v \n", err)
	}

	//Act
	resp, err := Handler(context.TODO())
	if err != nil {
		t.Errorf("\n Handler error: %+v \n", err)
	} else {
		fmt.Printf("\n resp: %+v \n", resp)
	}

}
