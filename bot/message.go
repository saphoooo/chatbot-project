package main

import "time"

// Message define the content of a bot answer
type Message struct {
	Name      string    `json:"name,omitempty"`
	Content   string    `json:"content,omitempty"`
	AvatarURL string    `json:"avatarURL,omitempty"`
	When      time.Time `json:"when,omitempty"`
	Bot       bool      `json:"bot,omitempty"`
}
