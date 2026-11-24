package storage

import (
	"encoding/json"

	"errors"
	"myproject/finance-tracker/finance"
	"os"
	"time"
)

type transactionDTO struct {
	ID              int
	Amount          float64
	Category        string
	TransactionType string
	Date            time.Time
	Comment         string
}

type JSONStorage struct {
	filePath string
}

func NewJSONStorage(filePath string) *JSONStorage {
	jsonStorage := &JSONStorage{
		filePath: filePath,
	}
	return jsonStorage
}

func (s *JSONStorage) Save(transactions []finance.Transaction) (err error) {
	dtos := make([]transactionDTO, 0)

	for _, transaction := range transactions {
		dto := transactionDTO{
			ID:              transaction.GetID(),
			Amount:          transaction.GetAmount(),
			Category:        transaction.GetCategory(),
			TransactionType: transaction.GetTransactionType(),
			Date:            transaction.GetDate(),
			Comment:         transaction.GetComment(),
		}

		dtos = append(dtos, dto)

	}
	data, err := json.Marshal(dtos)
	if err != nil {
		return err
	}
	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := file.Close()

		if err == nil {
			err = closeErr
		}

	}()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *JSONStorage) Load() (transactions []finance.Transaction, err error) {

	emptyTransactions := []finance.Transaction{}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return emptyTransactions, nil
		}
		return nil, err

	}

	var dtos []transactionDTO
	err = json.Unmarshal(data, &dtos)
	if err != nil {
		return nil, err
	}

	transactions = make([]finance.Transaction, 0, len(dtos))

	for _, dto := range dtos {
		transaction, err := finance.NewTransaction(
			dto.ID,
			dto.Amount,
			dto.Category,
			dto.TransactionType,
			dto.Date,
			dto.Comment,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, *transaction)
	}
	return transactions, nil

}
