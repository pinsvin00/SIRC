package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


var irc IRC

func main() {
	r := mux.NewRouter();

	irc.counter = 0
	irc.nicknames = []string{}
	irc.guid_nickname = make(map[string]string)
	irc.guid_color = make(map[string]string)
	irc.message_queue = make(map[string][]Message)
	irc.subscribers = 0

	r.HandleFunc("/get_credentials", send_credentials)
	r.HandleFunc("/subscribe", subscribe)
	r.HandleFunc("/send_message", message_handler)
	http.Handle("/", r);
	fmt.Print("Started listening on 9000")
	http.ListenAndServe(":9000", r)
}
