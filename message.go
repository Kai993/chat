package main

import "time"

// message : 1つのメッセージを表す
type message struct {
	Name    string
	Message string
	When    time.Time
}
