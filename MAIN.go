package main

import (
    "html/template"
    "net/http"
    "log"
    "encoding/json"
    "strings"
)

const get_messages_path, index_path string = "/get_messages", "/"
var valid_paths = [...]string {get_messages_path, index_path}

var messages string = ""

func check_path(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        for _, b := range valid_paths {
            if b == r.URL.Path {
                fn(w, r)
                return
            }
        }

        http.ServeFile(w, r, "404.html")
    }
}

func index(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, _ := template.ParseFiles("index.html")
        t.Execute(w, nil)
    } else if r.Method == "POST" {
        r.ParseForm()
        messages = strings.TrimSpace(messages + "\n" + r.Form["message"][0])
    }
}

func get_messages(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(messages)
}

func main() {
    http.HandleFunc(index_path, check_path(index)) // setting router rule
    http.HandleFunc(get_messages_path, check_path(get_messages))

    err := http.ListenAndServe("0.0.0.0:8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

