package models

import "time"

type File struct {
	Id              int64
	Name            string
	Size            int64
	CreatedTime     time.Time
	LastUpdatedTime time.Time
	IsDir           bool
}
