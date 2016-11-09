// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

var server *socketio.Server

func main() {
	var err error

	server, err = socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		so.Join("chat")
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.HandleFunc("/webhook", WebHook)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func WebHook(w http.ResponseWriter, r *http.Request)  {
	act := r.URL.Query().Get("act")
	text := r.URL.Query().Get("txt")

	if len(text) == 0 {
		fmt.Fprintln(w, "Request ignored, missing field [txt]!")
		return
	}
	if len(act) == 0 {
		fmt.Fprintln(w, "Request ignored, missing field [act]!")
		return
	}

	server.BroadcastTo("chat", act, text)
	fmt.Fprintf(w, "Message has been sent for action [%v]!\n", act)
}