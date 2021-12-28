package repo

import (
	"bank-test/entity"
)

type CustomerRepo interface {
	GetAll() []entity.Customer
}
