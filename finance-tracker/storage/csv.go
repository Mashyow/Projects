package storage

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"myproject/finance-tracker/finance"
	"os"
	"strconv"
	"time"
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

func (s *CSVStorage) Save(transactions []finance.Transaction) (err error) {

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

func (s *CSVStorage) Load() (transactions []finance.Transaction, err error) {

	emptyTransactions := []finance.Transaction{}
	file, err := os.Open(s.filePath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return emptyTransactions, nil
		}
		return nil, err
	}
	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}

	}()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {

		return nil, err
	}
	_ = header

	transactions = make([]finance.Transaction, 0)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}
		if len(record) != 6 {
			return nil, fmt.Errorf("invalid CSV record: expected 6 columns, got %d", len(record))
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err

		}

		amount, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}

		date, err := time.Parse("2006-01-02", record[4])
		if err != nil {
			return nil, err
		}

		transaction, err := finance.NewTransaction(
			id,
			amount,
			record[2],
			record[3],
			date,
			record[5],
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, *transaction)

	}
	return transactions, nil

}

var _ Storage = (*CSVStorage)(nil)
