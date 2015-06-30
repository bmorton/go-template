package db

import "time"

type UserRecord struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r TeamRecord) TableName() string {
	return "users"
}
