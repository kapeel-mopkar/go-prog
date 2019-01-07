package DAL

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/sample")
	if err != nil {
		log.Fatal("Database could not be connected")
	}
	return db
}

type Book struct {
	ID     int
	Name   string `json:"bookname,omitempty"`
	Isbn   string `json:"isbn"`
	Author string `json:"-"` // Not sent in JSON response
	Pages  int    `json:"number_of_pages"`
}

func GetBooks() []Book {
	var books []Book
	db := connect()
	defer db.Close()
	//Different ways for querying
	//1) Query : return multiple records
	//2) QueryRow : return single record

	rows, err := db.Query("select id, name, isbn, author, pages from books")
	if err != nil {
		log.Fatal("Books could not be retrieved")
	}

	//Iterate through rows and populate Books
	var book Book
	for rows.Next() {
		err1 := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Isbn, &book.Pages)
		if err != nil {
			log.Fatal("Book data could not be mapped", err1)
		}
		//Append book object to book slice
		books = append(books, book)
	}
	return books
}

func GetBookByID() Book {
	return
}

func GetTempBooks() []Book {
	var books []Book
	books = append(books, Book{101, "Origins", "IND001", "Dan Brown", 600},
		Book{102, "GOLang", "IND002", "Ajay Sharma", 200},
		Book{103, "MySQL", "IND003", "Mr. Frown", 300},
		Book{104, "MONGODB", "IND004", "Daniel Pearl", 550},
	)
	return books
}
