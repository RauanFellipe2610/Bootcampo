package main

import (
    "net/http"
)

func main() {
    // Serve o arquivo HTML diretamente
    http.Handle("/", http.FileServer(http.Dir("./static")))

    // Inicia o servidor na porta 8080
    http.ListenAndServe(":8080", nil)
}
