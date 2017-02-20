// TODO: Keep list of all sockets & send out msgs to them when received from 1
package main

import (
    "github.com/gorilla/websocket" // go get github.com/gorilla/websocket
    "html/template"
    "net/http"
    "log"
    //"encoding/json"
    //"strings"
    "fmt"
)

type ConnectionData struct {
   connection *websocket.Conn
}

const index_path, socket_path = "/", "/websocket"
var valid_paths = [...]string {index_path, socket_path}
var connections map[ConnectionData]bool

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

    var conn_obj = ConnectionData{conn}
    connections[conn_obj] = true

    defer delete(connections, conn_obj);
    defer conn.Close()

    for{
        _, msg, err := conn.ReadMessage() // Ignore msgType
        if err != nil {
            fmt.Println(err)
            return
        }

        var message = string(msg[:])

        for conn_data, _ := range connections{
            err = conn_data.connection.WriteMessage(websocket.TextMessage, []byte(message))

            if err != nil {
                fmt.Println(err)
                return
            }
        }
    }
}

func main() {
    connections = make(map[ConnectionData]bool)

    http.HandleFunc(index_path, check_path(index)) // setting router rule
    http.HandleFunc("/websocket", check_path(handle_socket))

    err := http.ListenAndServe("0.0.0.0:8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

