package main

import "time"

func NewTicker() *time.Ticker {
	return NewTicker(time.Second / 60)
}
