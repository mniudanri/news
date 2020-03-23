package main

import(
    "news-app/handlers"
    "net/http"
)

func init() {
    // Custom env file
}

func main(){
    http.HandleFunc("/news", handlers.HandleNews)

    http.ListenAndServe(":3000", nil)
}