package repo

import "bank-test/dto"

type LogRepo interface {
	Create(log dto.Log) error
}
