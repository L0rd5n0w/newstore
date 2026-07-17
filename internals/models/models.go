package models

import "time"

type Books struct {
	id 				int
	Title			string
	Author			string
	Description		string
	CreatedAt		time.Time
}