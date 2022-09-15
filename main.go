package main

import (
	"flag"
)

func main() {
	var host string
	var port string
	var database string
	flag.StringVar(&host, "h", "127.0.0.1", "hostname")
	flag.StringVar(&port, "p", "2233", "listening port")
	flag.StringVar(&database, "d", "database/lunyu.db", "quotes database")
	flag.Parse()
	db = OpenDB(database)

	Serve(host, port)

	CloseDB()
}
