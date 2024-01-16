package models

import (
	"fmt"
)

type Message struct {
	Content string
	Sender  *Client
	Time    string
}

func (m *Message) ToString() string {
	return fmt.Sprintf("[%v][%v]:%v", m.Time, m.Sender.Name, m.Content)
}
