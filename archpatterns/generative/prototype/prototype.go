package main

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Id        string
	Text      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewMessage(text string) *Message {
	return &Message{
		Id:        uuid.New().String(),
		Text:      text,
		CreatedAt: time.Now(),
	}
}

func (m *Message) Clone() *Message {
	return &Message{
		Id:        m.Id,
		Text:      m.Text,
		CreatedAt: m.CreatedAt,
		UpdatedAt: newLink(m.UpdatedAt),
	}
}
