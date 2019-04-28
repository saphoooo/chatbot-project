package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var chataddr = flag.String("chataddr", "localhost:8080", "http chat service address, et 127.0.0.0:8080")
var port = flag.String("port", ":9090", "bot service port")

func msgListener(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("read:", err)
	}
	log.Println(string(body))
	fallback()
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	rtr := mux.NewRouter()
	rtr.HandleFunc("/bot", msgListener).Methods("POST")
	http.Handle("/", rtr)
	log.Println("Start bot service on", *port)
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
