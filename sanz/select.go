package sanz

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Select(sanz string) {
	db, err := sql.Open("mysql", "root:sunlove8421!@tcp(localhost:3306)/taeyangyu")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var no int
	var book_name string
	rows, err := db.Query("SELECT no, book_name FROM book where no = ?", sanz)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)

	for rows.Next() {
		err := rows.Scan(&no, &book_name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(book_name)
	}

}
