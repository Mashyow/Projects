package wallet

import (
	"fmt"
	"myproject/finance-tracker/finance"
)

type Wallet struct {
	transactions []finance.Transaction
}

func (w *Wallet) AddTransaction(transaction finance.Transaction) {
	w.transactions = append(w.transactions, transaction)
}

func (w Wallet) GetTransactions() []finance.Transaction {
	return w.transactions
}

func (w Wallet) GetBalance() float64 {
	var balance float64 = 0
	for _, transaction := range w.transactions {
		amount := transaction.GetAmount()
		transactionType := transaction.GetTransactionType()

		if transactionType == "income" {
			balance += amount
		}
		if transactionType == "expense" {
			balance -= amount
		}
	}
	return balance
}

func (w Wallet) GetExpensesByCategory() map[string]float64 {
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
			return nil
		}

	}
	return fmt.Errorf("Транзакция не найдена")
}
