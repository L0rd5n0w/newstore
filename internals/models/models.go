package models

import (
	"time"
)

type Books struct {
	ID 				int
	Title			string
	Author			string
	Description		string
	CreatedAt		time.Time
}

type Users struct {
	FirstName			string
	LastName			string
	Email				string
	HashedPassword		[]byte
	Created				time.Time
}

