package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var chataddr = flag.String("chataddr", "localhost:8080", "http chat service address")
var port = flag.String("port", ":9090", "bot service port")
var botAvatar = flag.String("avatar", "white", "bot avatar")
var debug = flag.Bool("debug", false, "activate debug mode")

func stringInSlice(s string, list []string) bool {
	for _, avatar := range list {
		if avatar == s {
			return true
		}
	}
	return false
}

func msgListener(w http.ResponseWriter, r *http.Request) {
	var msg *Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Fatal(err)
	}
	switch msg.Content {
	case "Savez-vous ce qui s'est passé ce soir ?", "Que s'est-il passé ?", "Avez-vous vu quelque chose ?":
		msg.replyer(os.Getenv("QUOI"))
	case "Comment connaissez-vous Violette", "Quelles sont vos relations avec Violette":
		msg.replyer(os.Getenv("COMMENT"))
	case "Violette avait-elle des ennemis", "Violette était-elle fâchée avec quelqu'un", "Quelqu'un aurait pû lui vouloir du mal":
		msg.replyer(os.Getenv("QUI"))
	default:
		fallback()
	}
}

func main() {
	avatarList := []string{"blanc", "reglisse", "moutarde", "olive", "rose", "violette"}
	flag.Parse()
	if !stringInSlice(strings.ToLower(*botAvatar), avatarList) {
		log.Fatalf("avatar should be one of the following: %v\n", avatarList)
	}
	log.SetFlags(0)

	generateSomeLogs()

	rtr := mux.NewRouter()
	rtr.HandleFunc("/bot", msgListener).Methods("POST")
	http.Handle("/", rtr)
	if *debug {
		log.Println("Start bot service on", *port)
	}
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
