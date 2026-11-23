package main

import (
	"flag"
	"fmt"

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

}
