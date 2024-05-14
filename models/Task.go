package models

import "time"

type Task struct {
	Id          int64
	Title       string
	Description string
	StartTime   time.Time
}
