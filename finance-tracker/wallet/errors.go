package wallet

import "fmt"

type TransactionNotFoundError struct {
	ID int
}

func (e TransactionNotFoundError) Error() string {

	err := fmt.Sprintf("Транзакция с ID %d не найдена", e.ID)

	return err
}
