package airtable

import (
	"time"
)

type Account struct {
	ApiKey  string
	BaseId  string
	BaseUrl string
}

type Options struct {
	Filter     string
	Sort       []map[string]string
	MaxRecords int
	View       string
}

type records struct {
	Records []Record `json:"records"`
}

type Record struct {
	CreatedTime *time.Time             `json:"createdTime,omitempty"`
	ID          string                 `json:"id"`
	Fields      map[string]interface{} `json:"fields"`
}
