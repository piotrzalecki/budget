package repository

import "github.com/piotrzalecki/budget/internal/models"

type DatabaseRepo interface {
	CreateTransactionCategory(tcm models.TransactionCategory) error
	AllTransactionCategories() ([]models.TransactionCategory, error)
	DeleteTransactionCategory(id int) error
	GetTransactionCategoryById(id int) (models.TransactionCategory, error)
	UpdateTransactionCategory(tcm models.TransactionCategory) error
}
