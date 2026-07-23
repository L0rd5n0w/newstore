package models

import (
	"time"
	"github.com/google/uuid"
	
)

type Books struct {
	ID 				int
	Title			string
	Author			string
	Description		string
	CreatedAt		time.Time
}

type Users struct {
	Id					uuid.UUID
	FirstName			string
	LastName			string
	Email				string
	HashedPassword		[]byte
	Created				time.Time
}
