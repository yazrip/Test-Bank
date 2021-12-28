package service

import "bank-test/dto"

type LogService interface {
	Create(log dto.Log) error
}
