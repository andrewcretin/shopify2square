package httpClient

import (
	"fmt"
	"github.com/andrewcretin/shopify2square/httpClient/models"
	"github.com/andrewcretin/shopify2square/test"
	"testing"
)

func TestSearchOrders(t *testing.T) {

	tests := []struct {
		name string
		req  models.SearchOrderReq
	}{
		{
			name: "search-test",
			req: models.SearchOrderReq{
				LocationIds: []string{"G34QEDDSED5JQ"},
			},
		},
	}

	err := test.SetupTestEnvironment(t)
	if err != nil {
		t.Errorf("\n setupTestEnvironment error: %+v \n", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := SearchOrders(tt.req)
			if err != nil {
				t.Errorf("SearchOrders test: %s error: %+v", tt.name, err)
			} else {
				fmt.Printf("\n SearchOrders resp: %+v \n", resp)
			}
		})
	}
}
