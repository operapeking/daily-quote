package main

type Quote struct {
	Uid     int    `json:"uid"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
