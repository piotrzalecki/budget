package dbrepo

import (
	"context"
	"time"

	"github.com/piotrzalecki/budget/internal/models"
)

func (m *postgresDBRepo) CreateTransactionCategory(tcm models.TransactionCategory) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "INSERT INTO transactions_categories (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4)"

	_, err := m.DB.ExecContext(ctx, stmt,
		tcm.Name,
		tcm.Description,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) AllTransactionCategories() ([]models.TransactionCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, name, description, created_at, updated_at FROM transactions_categories ORDER BY name ASC"

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tcats []models.TransactionCategory

	for rows.Next() {
		var i models.TransactionCategory
		err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.Description,
			&i.UpdatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		tcats = append(tcats, i)
	}

	if err = rows.Err(); err != nil {
		return tcats, err
	}

	return tcats, nil
}

func (m *postgresDBRepo) DeleteTransactionCategory(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM transactions_categories WHERE id = $1"

	_, err := m.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil

}

func (m *postgresDBRepo) GetTransactionCategoryById(id int) (models.TransactionCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, name, description FROM transactions_categories WHERE id=$1"

	var cat models.TransactionCategory

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&cat.Id,
		&cat.Name,
		&cat.Description,
	)

	if err != nil {
		return cat, err
	}

	return cat, nil

}

func (m *postgresDBRepo) UpdateTransactionCategory(tcm models.TransactionCategory) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "UPDATE transactions_categories SET name=$1, description=$2, updated_at=$3 WHERE id=$4"

	_, err := m.DB.ExecContext(ctx, stmt,
		tcm.Name,
		tcm.Description,
		time.Now(),
		tcm.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
