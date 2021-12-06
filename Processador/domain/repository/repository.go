package repository

type TransactionRepository interface {
	Insert(id, account string, amount float64, status, errorMessage string) error
}
