package model

import "time"

type TimeCal struct {
	Interval int
	IsUpdate bool
	Next     time.Time
}