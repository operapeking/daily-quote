package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(file string) *sql.DB {
	db, err := sql.Open("sqlite3", "./quotes.db")
	if err != nil {
		panic(err)
	}
	return db
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func CreateTable() {
	table := `
    CREATE TABLE IF NOT EXISTS quote (
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT NULL UNIQUE,
        author TEXT
    );
    `
	_, err := db.Exec(table)
	if err != nil {
		panic(err)
	}
}

func ListAll() {
	rows, err := db.Query("SELECT * FROM quote")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var uid int
		var content string
		var author string
		err = rows.Scan(&uid, &content, &author)
		if err != nil {
			panic(err)
		}
		fmt.Println(uid, content, author)
	}
}

func QueryById(uid int) Quote {
	rows, err := db.Query(fmt.Sprintf("SELECT content, author FROM quote WHERE \"uid\"=\"%d\"", uid))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var res Quote
	err = rows.Scan(&res.Content, &res.Author)
	if err != nil {
		panic(err)
	}
	res.Uid = uid
	return res
}

func QueryByContent(content string) Quote {
	rows, err := db.Query(fmt.Sprintf("SELECT uid, author FROM quote WHERE \"content\"=\"%s\"", content))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var res Quote
	err = rows.Scan(&res.Uid, &res.Author)
	if err != nil {
		panic(err)
	}
	res.Content = content
	return res
}

func QueryByAuthor(author string) []Quote {

	rows, err := db.Query(fmt.Sprintf("SELECT uid, content FROM quote WHERE \"author\"=\"%s\"", author))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var res []Quote
	for rows.Next() {
		var now Quote
		err = rows.Scan(&now.Uid, &now.Content)
		if err != nil {
			panic(err)
		}
		now.Author = author
		res = append(res, now)
	}
	return res
}

func QueryByRandom() Quote {
	rows, err := db.Query("SELECT uid, content, author FROM quote ORDER BY RANDOM() limit 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var res Quote
	for rows.Next() {
		err = rows.Scan(&res.Uid, &res.Content, &res.Author)
		if err != nil {
			panic(err)
		}
	}
	return res
}

func Insert(quote Quote) (int64, error) {
	stmt, err := db.Prepare("INSERT OR IGNORE INTO quote(content, author) values(?, ?)")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(quote.Content, quote.Author)
	if err != nil {
		panic(err)
	}
	return res.LastInsertId()
}

func DeleteById(uid int) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM quote WHERE uid=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(uid)
	if err != nil {
		panic(err)
	}
	return res.RowsAffected()
}

func DeleteByContent(content string) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM quote WHERE content=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(content)
	if err != nil {
		panic(err)
	}
	return res.RowsAffected()
}

func DeleteByAuthor(author string) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM quote WHERE author=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(author)
	if err != nil {
		panic(err)
	}
	return res.RowsAffected()
}
