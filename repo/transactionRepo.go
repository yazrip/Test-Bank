package repo

import "bank-test/entity"

type TransactionRepo interface {
	Create(transaction entity.Transaction) error
	GetAll() []entity.Transaction
}
