package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func (m *Message) replyer(s string) {
	u := url.URL{Scheme: "ws", Host: *chataddr, Path: "/room"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	messg := newMsg()
	messg.Content = s
	msg, err := json.Marshal(messg)
	if err != nil {
		log.Println(err)
	}
	err = c.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("write:", err)
	}
}

func newMsg() *Message {
	avatar := strings.ToLower(*botAvatar)
	return &Message{Name: strings.Title(avatar), Bot: true, AvatarURL: "/avatars/" + avatar + ".png"}
}

func fallback() {
	m := Message{}
	var speech = []string{
		"Je n'ai pas compris. Pouvez-vous répéter ?",
		"J'ai raté ce que vous avez dit. Qu'est-ce que c'était ?",
		"Désolé, pourriez-vous répéter ?",
		"Désolé, pouvez-vous répéter ?",
		"Pouvez-vous répéter ?",
		"Désolé, je n'ai pas compris. Pouvez-vous reformuler ?",
		"Désolé, c'était quoi ?",
		"Encore une fois ?",
		"C'était quoi ?",
		"Dites-le encore une fois ?",
		"Je n'ai pas compris. Pouvez-vous répéter ?",
		"J'ai raté ça, pouvez-vous le répèter ?",
	}
	rand.Seed(time.Now().UTC().UnixNano())
	repl := speech[rand.Intn(len(speech))]
	m.replyer(repl)
}
