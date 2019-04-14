package main

import (
        "fmt"
        "net/http"
        "./routes"
        "github.com/gorilla/mux"
        )

func main() {
    r := mux.NewRouter()
    // r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))    
    
    routes.ApiRoutes(r)
    routes.NuxtjsRoutes(r)

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", r)
}