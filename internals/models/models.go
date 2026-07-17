package models

import "time"

type Books struct {
	ID 				int
	Title			string
	Author			string
	Description		string
	CreatedAt		time.Time
}