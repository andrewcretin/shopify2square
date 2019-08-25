package models

import "time"

type SearchOrderReq struct {
	ReturnEntries bool              `json:"return_entries,omitempty"`
	Limit         int               `json:"limit,omitempty"`
	LocationIds   []string          `json:"location_ids,omitempty"`
	Query         *SearchOrderQuery `json:"query,omitempty"`
	Cursor        *string           `json:"cursor,omitempty"`
}

type SearchOrderQuery struct {
	Filter *SearchFilter `json:"filter,omitempty"`
	Sort   *SearchSort   `json:"sort,omitempty"`
}

type SearchSort struct {
	SortField *string `json:"sort_field,omitempty"`
	SortOrder *string `json:"sort_order,omitempty"`
}

type SearchFilter struct {
	StateFilter       *SearchStateFilter       `json:"state_filter,omitempty"`
	DateTimeFilter    *SearchDateTimeFilter    `json:"date_time_filter,omitempty"`
	FulfillmentFilter *SearchFulfillmentFilter `json:"fulfillment_filter,omitempty"`
	SourceFilter      *SearchSourceFilter      `json:"source_filter,omitempty"`
	CustomerFilter    *SearchCustomerFilter    `json:"customer_filter,omitempty"`
}

type SearchDateTimeFilter struct {
	CreatedAt TimeRange `json:"created_at,omitempty"`
	UpdatedAt TimeRange `json:"updated_at,omitempty"`
	ClosedAt  TimeRange `json:"closed_at,omitempty"`
}

type SearchStateFilter struct {
	States []string `json:"states,omitempty"`
}

type SearchFulfillmentFilter struct {
	FulfillmentTypes  []string `json:"fulfillment_types,omitempty"`
	FulfillmentStates []string `json:"fulfillment_states,omitempty"`
}

type SearchSourceFilter struct {
	SourceNames []string `json:"source_names,omitempty"`
}

type SearchCustomerFilter struct {
	CustomerIds []string `json:"customer_ids,omitempty"` // max 10
}

type TimeRange struct {
	StartAt time.Time `json:"start_at,omitempty"`
	EndAt   time.Time `json:"end_at,omitempty"`
}
