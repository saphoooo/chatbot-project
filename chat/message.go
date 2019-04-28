package main

import "time"

type message struct {
	Name      string    `json:"name,omitempty"`
	Content   string    `json:"content,omitempty"`
	AvatarURL string    `json:"avatarURL,omitempty"`
	When      time.Time `json:"when,omitempty"`
	Bot       bool      `json:"bot,omitempty"`
}
