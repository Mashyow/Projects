package main

import (
	"bufio"
	"flag"
	"fmt"
	"myproject/finance-tracker/finance"
	"myproject/finance-tracker/storage"
	"myproject/finance-tracker/wallet"
	"os"
	"strconv"
	"time"
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

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("1. Добавить транзакцию")
		fmt.Println("2. Показать транзакции")
		fmt.Println("3. Статистика")
		fmt.Println("4. Удалить транзакцию")
		fmt.Println("5. Выход")

		fmt.Println("Выберите действие")
		if !scanner.Scan() {
			return
		}
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Некоректный ввод")
			continue
		}

		switch choice {

		case 1:

			fmt.Println("Введите ID:")
			if !scanner.Scan() {
				return
			}
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Некоректный ID")
				continue
			}

			fmt.Println("Введите сумму:")
			if !scanner.Scan() {
				return
			}
			amount, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("Некоректная сумма")
				continue
			}
			fmt.Println("ID", id, "Сумма", amount)

			fmt.Println("Введите категорию транзакции:")

			if !scanner.Scan() {
				return
			}
			category := scanner.Text()

			fmt.Println("Введите тип транзакции")
			fmt.Println("Доход", finance.TypeIncome)
			fmt.Println("Расход", finance.TypeExpense)
			if !scanner.Scan() {
				return
			}
			transactionType := scanner.Text()
			fmt.Println("Категория", category, "Тип транзакции", transactionType)

			fmt.Println("Введите дату в формате: ГГГГ-ММ-ДД:")
			if !scanner.Scan() {
				return
			}

			dateInput := scanner.Text()

			date, err := time.Parse("2006-01-02", dateInput)
			if err != nil {
				fmt.Println("Некорректная дата")
				continue
			}

			fmt.Println("Введите коментарий:")
			if !scanner.Scan() {
				return
			}
			comment := scanner.Text()

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

			fmt.Println("Введите ID транзакции:")
			if !scanner.Scan() {
				return
			}
			idTransaktion, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Некоректный ввод")
				continue
			}
			err = w.RemoveTransaction(idTransaktion)

			if err != nil {
				fmt.Println("Ошибка удаления", err)
				continue
			}
			fmt.Println("Транзакция удалена")

		case 5:
			err = w.Save()
			if err != nil {
				fmt.Println("Ошибка сохранения", err)
				return
			}

			fmt.Println("Данные сохранены. До свидания!")
			return

		default:
			fmt.Println("Неизвестный пункт меню")
		}

	}

}
