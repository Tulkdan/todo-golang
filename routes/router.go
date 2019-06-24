package routes

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/tulkdan/todo/logger"
)

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    var handler http.Handler
    handler = route.HandlerFunc
    handler = logger.Logger(handler, route.Name)

    router.
      Methods(route.Method).
      Path("/api" + route.Pattern).
      Name(route.Name).
      HandlerFunc(route.HandlerFunc)
  }

  return router
}