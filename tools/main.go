package main

import (
	"log"

	"github.com/Luxurioust/excelize"
)

func main() {
	f, err := excelize.OpenFile("./远程成员库.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}
