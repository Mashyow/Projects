package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"time"

	"myproject/finance-tracker/finance"
	"myproject/finance-tracker/storage"
	"myproject/finance-tracker/wallet"
)

func main() {
	var store storage.Storage
	format := flag.String("format", "json", "Формат хранения данных")
	flag.Parse()

	switch *format {
	case "json":
		store = storage.NewJSONStorage("transactions.json")

	case "csv":
		store = storage.NewCSVStorage("transactions.csv")

	default:
		fmt.Println("Неизвестный формат ")
		return
	}
	w, err := wallet.NewWallet(store)

	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}
	fmt.Println("Загружено транзакций:", len(w.GetTransactions()))

	for {
		fmt.Println("1. Добавить транзакцию")
		fmt.Println("2. Показать транзакции")
		fmt.Println("3. Статистика")
		fmt.Println("4. Удалить транзакцию")
		fmt.Println("5. Выход")

		var choice int
		fmt.Println("Выберите действие")
		_, err := fmt.Scan(&choice)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			} else {
				fmt.Println("Некоректный ввод")
				continue
			}
		}

		switch choice {

		case 1:
			var id int
			var amount float64
			fmt.Println("Введите ID:")
			fmt.Scan(&id)
			fmt.Println("Введите сумму:")
			fmt.Scan(&amount)
			fmt.Println("ID", id, "Сумма", amount)

			var category string
			var transactionType string
			fmt.Println("Введите категорию транзакии:")
			fmt.Scan(&category)
			fmt.Println("Введите тип транзакции")
			fmt.Scan(&transactionType)
			fmt.Println("Категория", category, "Тип транзакции", transactionType)

			var dateInput string
			fmt.Println("Введите дату в формате ГГГГ-ММ-ДД :")
			fmt.Scan(&dateInput)

			date, err := time.Parse("2006-01-02", dateInput)
			if err != nil {
				fmt.Println("Некорректная дата")
				continue
			}

			var comment string
			fmt.Println("Введите коментарий:")
			fmt.Scan(&comment)

			transaction, err := finance.NewTransaction(id, amount, category, transactionType, date, comment)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = w.AddTransaction(*transaction)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Транзакция добавлена ")

		case 2:
			if len(w.GetTransactions()) == 0 {
				fmt.Println("Транзакций пока нет")
			}
			for _, transaction := range w.GetTransactions() {
				fmt.Println(transaction.String())
			}

		case 3:
			balance := w.GetBalance()
			fmt.Println("Текущий баланс:", balance)

			expenses := w.GetExpensesByCategory()
			if len(expenses) == 0 {
				fmt.Println("Расходов пока нет")
				continue
			}
			for category, amount := range expenses {
				fmt.Println(category, amount)
			}

		case 4:
			var idTransaktion int
			fmt.Println("Введите ID транзакции:")
			fmt.Scan(&idTransaktion)
			err = w.RemoveTransaction(idTransaktion)
			if err != nil {
				fmt.Println("Ошибка удаления", err)
				continue
			}
			fmt.Println("Транзакция удалена")

		case 5:
			return

		default:
			fmt.Println("Неизвестный пункт меню")
		}

	}

}
