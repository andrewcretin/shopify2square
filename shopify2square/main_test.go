package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestHandler_test(t *testing.T) {

	err := setupTestEnvironment(t)
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

func setupTestEnvironment(t *testing.T) error {
	_ = os.Setenv("STAGE", "staging")
	env, err := parseEnvironment("staging")
	if err != nil {
		t.Errorf("\n parseEnvironment error: %+v \n", err)
	}
	// set the shopify key
	shopifyKey, ok := env["shopifyKey"]
	if ok {
		_ = os.Setenv("SHOPIFY_KEY", shopifyKey)
	} else {
		t.Errorf("\n must provide a valid 'shopifyKey' in config file \n")
	}
	// set the shopify key
	shopifyPassword, ok := env["shopifyPassword"]
	if ok {
		_ = os.Setenv("SHOPIFY_PASSWORD", shopifyPassword)
	} else {
		t.Errorf("\n must provide a valid 'shopifyPassword' in config file \n")
	}
	// set the square access token
	squareAccessToken, ok := env["squareAccessToken"]
	if ok {
		_ = os.Setenv("SQUARE_ACCESS_TOKEN", squareAccessToken)
	} else {
		t.Errorf("\n must provide a valid 'squareAccessToken' in config file \n")
	}

	return nil
}

func parseEnvironment(env string) (map[string]string, error) {

	resultsMap := map[string]string{}
	filePath := "../envConfig/config." + env + ".json"
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal(byteValue, &resultsMap)

	return resultsMap, nil

}
