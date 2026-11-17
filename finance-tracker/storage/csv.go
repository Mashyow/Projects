package storage

import (
	"encoding/csv"
	"myproject/finance-tracker/finance"
	"os"
	"strconv"
)

type CSVStorage struct {
	filePath string
}

func NewCSVStorage(filePath string) *CSVStorage {
	csvStorage := &CSVStorage{
		filePath: filePath,
	}
	return csvStorage
}

func (s *CSVStorage) Save(transactions []finance.Transaction) error {

	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	if err := writer.Write([]string{
		"ID",
		"Amount",
		"Category",
		"TransactionType",
		"Date",
		"Comment",
	}); err != nil {
		return err
	}

	for _, transaction := range transactions {
		row := []string{

			strconv.Itoa(transaction.GetID()),
			strconv.FormatFloat(transaction.GetAmount(), 'f', 2, 64),
			transaction.GetCategory(),
			transaction.GetTransactionType(),
			transaction.GetDate().Format("2006-01-02"),
			transaction.GetComment(),
		}
		if err := writer.Write(row); err != nil {
			return err

		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
