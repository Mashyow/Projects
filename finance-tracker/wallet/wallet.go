package wallet

import (
	"fmt"
	"myproject/finance-tracker/finance"
	"myproject/finance-tracker/storage"
)

func NewWallet(storage storage.Storage) (*Wallet, error) {
	wallet := &Wallet{
		storage: storage,
	}
	transactions, err := storage.Load()
	if err != nil {
		return nil, err
	}

	wallet.transactions = transactions
	return wallet, nil
}

func (w *Wallet) Save() error {
	err := w.storage.Save(w.transactions)

	if err != nil {

		return err
	}
	return nil
}

type Wallet struct {
	transactions []finance.Transaction

	storage storage.Storage
}

func (w *Wallet) AddTransaction(transaction finance.Transaction) error {

	for _, existing := range w.transactions {
		if existing.GetID() == transaction.GetID() {
			return fmt.Errorf("ID %d уже существует", transaction.GetID())
		}
	}

	w.transactions = append(w.transactions, transaction)
	err := w.Save()
	if err != nil {
		return err
	}
	return nil
}

func (w Wallet) GetTransactions() []finance.Transaction {

	return w.transactions
}

func (w *Wallet) GetBalance() float64 {
	var balance float64 = 0
	for _, transaction := range w.transactions {
		amount := transaction.GetAmount()
		transactionType := transaction.GetTransactionType()

		if transactionType == finance.TypeIncome {
			balance += amount
		}
		if transactionType == finance.TypeExpense {
			balance -= amount
		}
	}
	return balance
}

func (w *Wallet) GetExpensesByCategory() map[string]float64 {
	expenses := make(map[string]float64)

	for _, transaction := range w.transactions {
		transactionType := transaction.GetTransactionType()
		category := transaction.GetCategory()
		amount := transaction.GetAmount()
		if transactionType == "expense" {
			expenses[category] += amount
		}

	}
	return expenses
}

func (w *Wallet) RemoveTransaction(id int) error {

	for index, transaction := range w.transactions {
		ID := transaction.GetID()

		if ID == id {
			w.transactions = append(w.transactions[:index], w.transactions[index+1:]...)
			err := w.Save()
			if err != nil {
				return err
			}
			return nil
		}

	}
	return TransactionNotFoundError{ID: id}
}
