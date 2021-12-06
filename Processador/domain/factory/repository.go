package factory

import (
	"github.com/vellone-fabricio/imersao5-gateway/domain/repository"
)

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
