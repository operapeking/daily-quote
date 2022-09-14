package main

import (
	"flag"
)

func main() {
	var host string
	var port string
	flag.StringVar(&host, "h", "127.0.0.1", "hostname")
	flag.StringVar(&port, "p", "2233", "listening port")
	flag.Parse()
	db = OpenDB("./quotes.db")
	CreateTable()

	Insert(quote1)
	Insert(quote2)
	Insert(quote3)
	Serve(host, port)

	CloseDB()
}
