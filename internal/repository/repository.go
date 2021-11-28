package repository

import (
	"time"

	"github.com/piotrzalecki/budget/internal/models"
)

type DatabaseRepo interface {
	CreateTransactionCategory(tcm models.TransactionCategory) (int, error)
	AllTransactionCategories() ([]models.TransactionCategory, error)
	DeleteTransactionCategory(id int) error
	GetTransactionCategoryById(id int) (models.TransactionCategory, error)
	UpdateTransactionCategory(tcm models.TransactionCategory) error
	CreateRecurentTransaction(rt models.TransactionRecurence) (int, error)
	AllRecurentTransactions() ([]models.TransactionRecurence, error)
	DeleteRecurentTransaction(id int) error
	GetRecurentTransactionById(id int) (models.TransactionRecurence, error)
	UpdateRecurentTransaction(rt models.TransactionRecurence) error
	CreateTransactionType(tt models.TransactionType) (int, error)
	AllTransactionTypes() ([]models.TransactionType, error)
	DeleteTransactionType(id int) error
	GetTransactionTypeById(id int) (models.TransactionType, error)
	UpdateTransactionType(tt models.TransactionType) error
	CreateTransactionData(td models.TransactionData) (int, error)
	AllTransactionsData() ([]models.TransactionData, error)
	DeleteTransactionData(id int) error
	GetTransactionDataById(id int) (models.TransactionData, error)
	UpdateTransactionsData(td models.TransactionData) error
	GetAllActiveRecurentTransactions(to_date time.Time) ([]models.TransactionData, error)
	GetLatestBalanceQuote() (float32, error)
	GetSingleTransactionsForDates(from_date, to_date time.Time) ([]models.TransactionData, error)
	CreateTransactionLog(tl models.TransactionLog) (int, error)
	AllTransactionsLogs() ([]models.TransactionLog, error)
	GetTransactionLogById(id int) (models.TransactionLog, error)
	CreateAccountBalance(ab models.AccountBalance) (int, error)
	AllTransactionsLogsForDates(from_date, to_date time.Time) ([]models.TransactionLog, error)
}
