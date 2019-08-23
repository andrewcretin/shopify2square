package main

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestHandler_test(t *testing.T) {

	_ = os.Setenv("STAGE", "staging")

	//Act
	resp, err := Handler(context.TODO())
	if err != nil {
		t.Errorf("\n error: %+v \n", err)
	} else {
		fmt.Printf("\n resp: %+v \n", resp)
	}

}
