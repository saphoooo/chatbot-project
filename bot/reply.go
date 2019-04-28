package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func (m *Message) replyer(s string) {
	u := url.URL{Scheme: "ws", Host: *chataddr, Path: "/room"}
	//log.Printf("connecting to %s", u.String())
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
	return &Message{Name: "beepbop", Bot: true, AvatarURL: "/avatars/bot.png"}
}

func fallback() {
	m := Message{}
	var speech = []string{
		"I didn't get that. Can you say it again?",
		"I missed what you said. What was that?",
		"Sorry, could you say that again?",
		"Sorry, can you say that again?",
		"Can you say that again?",
		"Sorry, I didn't get that. Can you rephrase?",
		"Sorry, what was that?",
		"One more time?",
		"What was that?",
		"Say that one more time?",
		"I didn't get that. Can you repeat?",
		"I missed that, say that again?",
	}
	rand.Seed(time.Now().UTC().UnixNano())
	repl := speech[rand.Intn(len(speech))]
	m.replyer(repl)
}
