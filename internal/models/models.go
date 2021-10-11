package models

import "time"

type TemplateData struct {
	Data map[string]interface{}
}

type TransactionCategory struct {
	Id          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
