package main

import (
    "fmt"
    "html/template"
    "net/http"
    "log"
    "encoding/json"
    "strings"
)

var messages string = ""

func index(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        messages = strings.TrimSpace(messages + "\n" + r.Form["message"][0])
    }

    t, _ := template.ParseFiles("index.html")
    t.Execute(w, nil)
}

func get_messages(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(messages)
}

func main() {
    http.HandleFunc("/", index) // setting router rule
    http.HandleFunc("/get_messages", get_messages)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

