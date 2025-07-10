package entity

import "time"

type ToDo struct {
	Id        uint
	Name      string
	Priority  string
	StartTime time.Time
	EndTime   time.Time
}
