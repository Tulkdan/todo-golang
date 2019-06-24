package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/tulkdan/todo/routes"
)

func main() {
  router := routes.NewRouter()

  fmt.Println("Starting server...")
  log.Fatal(http.ListenAndServe(":8000", router))
}