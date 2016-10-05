package models

import (
	"fmt"
	"time"
)

type Score struct {
	Id       int64
	Value    int64
	Nickname string
	Inserted string
}

func (s Score) String() string {
	return fmt.Sprintf("%s: %s", s.Nickname, s.Value)
}

func NewScore(value int64, nickname string) Score {
	t := time.Now().Format(time.RFC822)
	return Score{
		Value:    value,
		Nickname: nickname,
		Inserted: t,
	}
}
