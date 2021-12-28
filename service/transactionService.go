package service

import "bank-test/entity"

type TransactionService interface {
	Create(transaction entity.Transaction) error
	GetAll() []entity.Transaction
}
