package main

import(
    "news-app/handlers"
    "net/http"
)

func main(){

    http.HandleFunc("/news", handlers.HandleNews)

    http.ListenAndServe(":3000", nil)
}