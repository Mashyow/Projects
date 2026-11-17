package storage

import "myproject/finance-tracker/finance"

type Storage interface {
	Save([]finance.Transaction) error

	Load() ([]finance.Transaction, error)
}
