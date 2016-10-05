package models

import (
	"time"
)

type Fasta struct {
	Id       int64 "db: auto_increment"
	Inserted time.Time
	Title    string
	Content  string
}

func NewFasta(title, content string) Fasta {
	return Fasta{
		Inserted: time.Now(),
		Title:    title,
		Content:  content,
	}
}
