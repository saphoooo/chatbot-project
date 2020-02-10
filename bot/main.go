package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("read:", err)
	}
	log.Println(string(body))
	fallback()
}

func main() {
	avatarList := []string{"white", "black", "mustard", "olive", "rose", "violet"}
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
