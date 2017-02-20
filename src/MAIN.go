// TODO: Keep list of all sockets & send out msgs to them when received from 1
package main

import (
    "github.com/gorilla/websocket" // go get github.com/gorilla/websocket
    "html/template"
    "net/http"
    "log"
    //"encoding/json"
    "strings"
    "fmt"
)

const index_path = "/"
var valid_paths = [...]string {index_path}

var messages string = ""

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

func check_path(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        for _, b := range valid_paths {
            if b == r.URL.Path {
                fn(w, r)
                return
            }
        }

        http.ServeFile(w, r, "html/404.html")
    }
}

func index(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, _ := template.ParseFiles("html/index.html")
        t.Execute(w, nil)
    }
}

func handle_socket(w http.ResponseWriter, r *http.Request){
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    for{
        _, msg, err := conn.ReadMessage() // Ignore msgType
        if err != nil {
            fmt.Println(err)
            return
        }

        var message = string(msg[:])
        messages = strings.TrimSpace(messages + "\n" + message)

        err = conn.WriteMessage(websocket.TextMessage, []byte(message)) // messages
        if err != nil {
            fmt.Println(err)
            return
        }
    }
}

func main() {
    http.HandleFunc(index_path, check_path(index)) // setting router rule
    http.HandleFunc("/websocket", handle_socket)

    err := http.ListenAndServe("0.0.0.0:8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

