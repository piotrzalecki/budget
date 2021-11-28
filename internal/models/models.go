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

type TransactionRecurence struct {
	Id          int
	Name        string
	Description string
	AddTime     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TransactionType struct {
	Id          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TransactionData struct {
	Id                   int
	Name                 string
	Description          string
	TransactionQuote     float32
	TransactionDate      time.Time
	TransactionType      TransactionType
	TransactionCategory  TransactionCategory
	TransactionRecurence TransactionRecurence
	RepeatUntil          time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type TransactionsData []TransactionData

type TransactionLog struct {
	Id               int
	TransactionData  TransactionData
	TransactionQuote float32
	TransactionDate  time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CreatedBy        User
	UpdateBy         User
}

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ActivityType struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ActivityLog struct {
	Id        int
	Type      ActivityType
	User      User
	CreatedAt time.Time
}

type AccountBalance struct {
	Id                 int
	Balance            float32
	BalanceTransaction TransactionLog
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
