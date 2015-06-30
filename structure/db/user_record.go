package db

import "time"

type UserRecord struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r UserRecord) TableName() string {
	return "users"
}
