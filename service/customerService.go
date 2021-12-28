package service

import "bank-test/entity"

type CustomerService interface {
	GetAll() []entity.Customer
}
