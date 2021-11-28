package dbrepo

import (
	"context"
	"time"

	"github.com/piotrzalecki/budget/internal/models"
)

// TRANSACTION CATEGORIES
func (m *postgresDBRepo) CreateTransactionCategory(tcm models.TransactionCategory) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "INSERT INTO transactions_categories (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4)  RETURNING id"
	var id int
	err := m.DB.QueryRowContext(ctx, stmt,
		tcm.Name,
		tcm.Description,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
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

	query := "SELECT id, name, description, created_at, updated_at FROM transactions_categories WHERE id=$1"

	var cat models.TransactionCategory

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&cat.Id,
		&cat.Name,
		&cat.Description,
		&cat.CreatedAt,
		&cat.UpdatedAt,
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

// RECURENT TRANSACTIONS
func (m *postgresDBRepo) CreateRecurentTransaction(tr models.TransactionRecurence) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "INSERT INTO transactions_recurence (name, description, addtime, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := m.DB.QueryRowContext(ctx, stmt,
		tr.Name,
		tr.Description,
		tr.AddTime,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *postgresDBRepo) AllRecurentTransactions() ([]models.TransactionRecurence, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, name, description, addtime, created_at, updated_at FROM transactions_recurence ORDER BY name ASC"

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trs []models.TransactionRecurence

	for rows.Next() {
		var i models.TransactionRecurence
		err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.Description,
			&i.AddTime,
			&i.UpdatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		trs = append(trs, i)
	}

	if err = rows.Err(); err != nil {
		return trs, err
	}

	return trs, nil
}

func (m *postgresDBRepo) DeleteRecurentTransaction(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM transactions_recurence WHERE id = $1"

	_, err := m.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil

}

func (m *postgresDBRepo) GetRecurentTransactionById(id int) (models.TransactionRecurence, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, name, description, addtime, created_at, updated_at FROM transactions_recurence WHERE id=$1"

	var tr models.TransactionRecurence

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&tr.Id,
		&tr.Name,
		&tr.Description,
		&tr.AddTime,
		&tr.CreatedAt,
		&tr.UpdatedAt,
	)

	if err != nil {
		return tr, err
	}

	return tr, nil

}

func (m *postgresDBRepo) UpdateRecurentTransaction(rt models.TransactionRecurence) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "UPDATE transactions_recurence SET name=$1, description=$2, addtime=$3, updated_at=$4 WHERE id=$5"

	_, err := m.DB.ExecContext(ctx, stmt,
		rt.Name,
		rt.Description,
		rt.AddTime,
		time.Now(),
		rt.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

// TRANSACTION TYPES
func (m *postgresDBRepo) CreateTransactionType(tt models.TransactionType) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "INSERT INTO trnsactions_types (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4)  RETURNING id"
	var id int
	err := m.DB.QueryRowContext(ctx, stmt,
		tt.Name,
		tt.Description,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *postgresDBRepo) AllTransactionTypes() ([]models.TransactionType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, name, description, created_at, updated_at FROM transactions_types ORDER BY name ASC"

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tts []models.TransactionType

	for rows.Next() {
		var i models.TransactionType
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

		tts = append(tts, i)
	}

	if err = rows.Err(); err != nil {
		return tts, err
	}

	return tts, nil
}

func (m *postgresDBRepo) DeleteTransactionType(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM transactions_types WHERE id = $1"

	_, err := m.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil

}

func (m *postgresDBRepo) GetTransactionTypeById(id int) (models.TransactionType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, name, description, created_at, updated_at FROM transactions_types WHERE id=$1"

	var tt models.TransactionType

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&tt.Id,
		&tt.Name,
		&tt.Description,
		&tt.CreatedAt,
		&tt.UpdatedAt,
	)

	if err != nil {
		return tt, err
	}

	return tt, nil

}

func (m *postgresDBRepo) UpdateTransactionType(tt models.TransactionType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "UPDATE transactions_types SET name=$1, description=$2, updated_at=$3 WHERE id=$4"

	_, err := m.DB.ExecContext(ctx, stmt,
		tt.Name,
		tt.Description,
		time.Now(),
		tt.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

// TRANSACTIONs DATA
func (m *postgresDBRepo) CreateTransactionData(td models.TransactionData) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO transactions_data (name, description, transaction_quote, transaction_date, transaction_type, transaction_category, transaction_recurence, repeat_until, created_at, updated_at) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)  RETURNING id`

	var id int
	err := m.DB.QueryRowContext(ctx, stmt,
		td.Name,
		td.Description,
		td.TransactionQuote,
		td.TransactionDate,
		td.TransactionType.Id,
		td.TransactionCategory.Id,
		td.TransactionRecurence.Id,
		td.RepeatUntil,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *postgresDBRepo) AllTransactionsData() ([]models.TransactionData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT td.id, td.name, td.transaction_quote, td.transaction_date, td.repeat_until,
	tc.name, tt.name, tr.name
	FROM transactions_data td 
	LEFT JOIN transactions_categories tc ON (td.transaction_category = tc.id)
	LEFT JOIN transactions_types tt ON (td.transaction_type = tt.id)
	LEFT JOIN transactions_recurence tr ON (td.transaction_recurence = tr.id)
	ORDER BY td.transaction_date ASC`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tds []models.TransactionData

	for rows.Next() {
		var i models.TransactionData
		err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.TransactionQuote,
			&i.TransactionDate,
			&i.RepeatUntil,
			&i.TransactionCategory.Name,
			&i.TransactionType.Name,
			&i.TransactionRecurence.Name,
		)
		if err != nil {
			return nil, err
		}

		tds = append(tds, i)
	}

	if err = rows.Err(); err != nil {
		return tds, err
	}

	return tds, nil
}

func (m *postgresDBRepo) DeleteTransactionData(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM transactions_data WHERE id = $1"

	_, err := m.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil

}

func (m *postgresDBRepo) GetTransactionDataById(id int) (models.TransactionData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT td.id, td.name, td.description, td.transaction_quote, td.transaction_date, td.repeat_until, td.created_at, td.updated_at,
	tc.id, tc.name, tc.description, tc.created_at, tc.updated_at,  tt.id, tt.name, tt.description,tt.created_at, tt.updated_at,
	 tr.id, tr.name, tr.description, tr.addtime, tr.created_at, tr.updated_at
	FROM transactions_data td
	LEFT JOIN transactions_categories tc ON (td.transaction_category = tc.id)
	LEFT JOIN transactions_types tt ON (td.transaction_type = tt.id)
	LEFT JOIN transactions_recurence tr ON (td.transaction_recurence = tr.id)
	WHERE td.id = $1
	ORDER BY td.transaction_date ASC`

	var td models.TransactionData

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&td.Id,
		&td.Name,
		&td.Description,
		&td.TransactionQuote,
		&td.TransactionDate,
		&td.RepeatUntil,
		&td.CreatedAt,
		&td.UpdatedAt,
		&td.TransactionCategory.Id,
		&td.TransactionCategory.Name,
		&td.TransactionCategory.Description,
		&td.TransactionCategory.CreatedAt,
		&td.TransactionCategory.UpdatedAt,
		&td.TransactionType.Id,
		&td.TransactionType.Name,
		&td.TransactionType.Description,
		&td.TransactionType.CreatedAt,
		&td.TransactionType.UpdatedAt,
		&td.TransactionRecurence.Id,
		&td.TransactionRecurence.Name,
		&td.TransactionRecurence.Description,
		&td.TransactionRecurence.AddTime,
		&td.TransactionRecurence.CreatedAt,
		&td.TransactionRecurence.UpdatedAt,
	)

	if err != nil {
		return td, err
	}

	return td, nil

}

func (m *postgresDBRepo) UpdateTransactionsData(td models.TransactionData) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "UPDATE transactions_data SET name=$1, description=$2, transaction_quote=$3, transaction_date=$4, transaction_type=$5, transaction_category=$6, transaction_recurence=$7, repeat_until=$8, updated_at=$9 WHERE id=$10"
	_, err := m.DB.ExecContext(ctx, stmt,
		td.Name,
		td.Description,
		td.TransactionQuote,
		td.TransactionDate,
		td.TransactionType.Id,
		td.TransactionCategory.Id,
		td.TransactionRecurence.Id,
		td.RepeatUntil,
		time.Now(),
		td.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

// Timeline queries

func (m *postgresDBRepo) GetAllActiveRecurentTransactions(to_date time.Time) ([]models.TransactionData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT td.id, td.name, td.transaction_quote, td.transaction_date, td.repeat_until,
	tt.name, tr.addtime
	FROM transactions_data td 
	LEFT JOIN transactions_types tt ON (td.transaction_type = tt.id)
	LEFT JOIN transactions_recurence tr ON (td.transaction_recurence = tr.id)
	WHERE td.repeat_until > $1 AND tt.name='RT'
	ORDER BY td.transaction_date ASC`

	rows, err := m.DB.QueryContext(ctx, query, to_date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tds []models.TransactionData

	for rows.Next() {
		var i models.TransactionData
		err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.TransactionQuote,
			&i.TransactionDate,
			&i.RepeatUntil,
			&i.TransactionType.Name,
			&i.TransactionRecurence.AddTime,
		)
		if err != nil {
			return nil, err
		}

		tds = append(tds, i)
	}

	if err = rows.Err(); err != nil {
		return tds, err
	}

	return tds, nil
}

func (m *postgresDBRepo) GetSingleTransactionsForDates(from_date, to_date time.Time) ([]models.TransactionData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT td.id, td.name, td.transaction_quote, td.transaction_date, tt.name
	FROM transactions_data td 
	LEFT JOIN transactions_types tt ON (td.transaction_type = tt.id)
	WHERE td.transaction_date > $1 AND td.transaction_date < $2 AND tt.name='ST'
	ORDER BY td.transaction_date ASC`

	rows, err := m.DB.QueryContext(ctx, query, from_date, to_date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tds []models.TransactionData

	for rows.Next() {
		var i models.TransactionData
		err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.TransactionQuote,
			&i.TransactionDate,
			&i.TransactionType.Name,
		)
		if err != nil {
			return nil, err
		}

		tds = append(tds, i)
	}

	if err = rows.Err(); err != nil {
		return tds, err
	}

	return tds, nil
}

// Account Balance queries

func (m *postgresDBRepo) GetLatestBalanceQuote() (float32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, balance FROM account_balance ORDER BY id DESC LIMIT 1;`

	var ab models.AccountBalance

	row := m.DB.QueryRowContext(ctx, query)
	err := row.Scan(
		&ab.Id,
		&ab.Balance,
	)

	if err != nil {
		return 0, err
	}

	return ab.Balance, nil

}

func (m *postgresDBRepo) CreateAccountBalance(ab models.AccountBalance) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO account_balance (balance, balance_transaction, created_at, updated_at) 
			VALUES ($1, $2, $3, $4)  RETURNING id`

	var id int
	err := m.DB.QueryRowContext(ctx, stmt,
		ab.Balance,
		ab.BalanceTransaction.Id,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil

}

// TRANSACTIONS LOG
func (m *postgresDBRepo) CreateTransactionLog(tl models.TransactionLog) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO transactions_logs (transaction_data, transaction_quote, transaction_date, created_at, updated_at, created_by, updated_by) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)  RETURNING id`

	var id int
	err := m.DB.QueryRowContext(ctx, stmt,
		tl.TransactionData.Id,
		tl.TransactionQuote,
		tl.TransactionDate,
		time.Now(),
		time.Now(),
		tl.CreatedBy.Id,
		tl.UpdateBy.Id,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *postgresDBRepo) AllTransactionsLogs() ([]models.TransactionLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT tl.id, tl.transaction_quote, tl.transaction_date, tl.created_at, tl.updated_at, tl.created_by, tl.updated_by
	FROM transactions_logs tl
	ORDER BY tl.transaction_date DESC LIMIT 50`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tl []models.TransactionLog

	for rows.Next() {
		var i models.TransactionLog
		err := rows.Scan(
			&i.Id,
			&i.TransactionQuote,
			&i.TransactionDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy.Id,
			&i.UpdateBy.Id,
		)
		if err != nil {
			return nil, err
		}

		tl = append(tl, i)
	}

	if err = rows.Err(); err != nil {
		return tl, err
	}

	return tl, nil
}

func (m *postgresDBRepo) AllTransactionsLogsForDates(from_date, to_date time.Time) ([]models.TransactionLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT tl.id, tl.transaction_data, tl.transaction_quote, tl.transaction_date, tl.created_at, tl.updated_at, tl.created_by, tl.updated_by
	FROM transactions_logs tl
	WHERE tl.transaction_date > $1 AND tl.transaction_date < $2
	ORDER BY tl.transaction_date DESC`

	rows, err := m.DB.QueryContext(ctx, query, from_date, to_date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tl []models.TransactionLog

	for rows.Next() {
		var i models.TransactionLog
		err := rows.Scan(
			&i.Id,
			&i.TransactionData.Id,
			&i.TransactionQuote,
			&i.TransactionDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy.Id,
			&i.UpdateBy.Id,
		)
		if err != nil {
			return nil, err
		}

		tl = append(tl, i)
	}

	if err = rows.Err(); err != nil {
		return tl, err
	}

	return tl, nil
}

// func (m *postgresDBRepo) DeleteTransactionData(id int) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	query := "DELETE FROM transactions_data WHERE id = $1"

// 	_, err := m.DB.ExecContext(ctx, query, id)

// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

func (m *postgresDBRepo) GetTransactionLogById(id int) (models.TransactionLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT tl.id, tl.transaction_quote, tl.transaction_date, tl.created_at, tl.updated_at, tl.created_by, tl.updated_by
	FROM transactions_logs tl 
	WHERE tl.id= $1
	ORDER BY tl.transaction_date ASC`

	var i models.TransactionLog

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&i.Id,
		&i.TransactionQuote,
		&i.TransactionDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy.Id,
		&i.UpdateBy.Id,
	)

	if err != nil {
		return i, err
	}

	return i, nil

}

// func (m *postgresDBRepo) UpdateTransactionsData(td models.TransactionData) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	stmt := "UPDATE transactions_data SET name=$1, description=$2, transaction_quote=$3, transaction_date=$4, transaction_type=$5, transaction_category=$6, transaction_recurence=$7, repeat_until=$8, updated_at=$9 WHERE id=$10"
// 	_, err := m.DB.ExecContext(ctx, stmt,
// 		td.Name,
// 		td.Description,
// 		td.TransactionQuote,
// 		td.TransactionDate,
// 		td.TransactionType.Id,
// 		td.TransactionCategory.Id,
// 		td.TransactionRecurence.Id,
// 		td.RepeatUntil,
// 		time.Now(),
// 		td.Id,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (m *postgresDBRepo) GetTransactionLogByTransactionDataIdFor(id int) (models.TransactionLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT tl.id, tl.transaction_quote, tl.transaction_date, tl.created_at, tl.updated_at, tl.created_by, tl.updated_by
	FROM transactions_logs tl 
	WHERE tl.transaction_data= $1
	ORDER BY tl.transaction_date ASC`

	var i models.TransactionLog

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&i.Id,
		&i.TransactionQuote,
		&i.TransactionDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy.Id,
		&i.UpdateBy.Id,
	)

	if err != nil {
		return i, err
	}

	return i, nil

}
