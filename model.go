package main

type Quote struct {
	Uid     int    `json:"uid"`
	Content string `json:"content"`
	From    string `json:"from"`
	Author  string `json:"author"`
}
