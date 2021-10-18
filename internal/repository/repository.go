package repository

import "github.com/piotrzalecki/budget/internal/models"

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
}
