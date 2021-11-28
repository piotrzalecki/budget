package models

func (td TransactionsData) Len() int {
	return len(td)
}

func (td TransactionsData) Less(i, j int) bool {
	return td[i].TransactionDate.Before(td[j].TransactionDate)
}

func (td TransactionsData) Swap(i, j int) {
	td[i], td[j] = td[j], td[i]
}
