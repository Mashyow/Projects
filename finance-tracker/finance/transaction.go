package finance

import (
	"fmt"
	"strings"
	"time"
)

const (
	TypeIncome  = "income"
	TypeExpense = "expense"
)

type Transaction struct {
	id              int       //номер операции
	amount          float64   //сумма операции
	category        string    //категория
	transactionType string    //доход или расход
	date            time.Time //дата операции
	comment         string    //коментарий
}

func (a Transaction) GetID() int {
	return a.id
}

func (a Transaction) GetAmount() float64 {
	return a.amount
}

func (a Transaction) GetCategory() string {
	return a.category
}
func (a Transaction) GetTransactionType() string {
	return a.transactionType
}
func (a Transaction) GetDate() time.Time {
	return a.date
}
func (a Transaction) GetComment() string {
	return a.comment
}

func (a *Transaction) SetAmount(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Сумма операции должна быть больше 0")
	}
	a.amount = amount
	return nil
}

func (a *Transaction) SetCategory(category string) error {
	if strings.TrimSpace(category) == "" {
		return fmt.Errorf("Строка не должна быть пустой")
	}
	a.category = strings.TrimSpace(category)
	return nil
}

func (a *Transaction) SetTransactionType(transactionType string) error {
	cleanType := strings.TrimSpace(transactionType)
	if cleanType != TypeIncome && cleanType != TypeExpense {
		return fmt.Errorf("доход или расход!!!")
	}
	a.transactionType = cleanType
	return nil
}
func (a *Transaction) SetComment(comment string) {
	a.comment = comment
}

func NewTransaction(
	id int,
	amount float64,
	category string,
	transactionType string,
	date time.Time,
	comment string,
) (*Transaction, error) {
	transaction := &Transaction{
		id:      id,
		date:    date,
		comment: comment,
	}
	if err := transaction.SetAmount(amount); err != nil {
		return nil, err
	}
	if err := transaction.SetCategory(category); err != nil {
		return nil, err
	}

	if err := transaction.SetTransactionType(transactionType); err != nil {
		return nil, err
	}

	return transaction, nil

}

func (a Transaction) String() string {
	return fmt.Sprintf("ID: %d | Сумма: %.2f | Категория: %s | Тип операции: %s | Дата: %v | Комментарий: %s",
		a.id,
		a.amount,
		a.category,
		a.transactionType,
		a.date,
		a.comment,
	)
}
