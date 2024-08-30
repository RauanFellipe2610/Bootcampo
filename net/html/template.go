package main

import (
    "html/template"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("template.html"))
    
    // Dados que serão passados para o template
    data := struct {
        Title   string
        Message string
    }{
        Title:   "Meu Site em Go",
        Message: "Olá, seja bem-vindo!",
    }
    
    tmpl.Execute(w, data)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
