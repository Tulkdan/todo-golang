package routes

import (
  "net/http"
  "github.com/tulkdan/todo/handlers"
)

type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    Name: "Index",
    Method: "GET",
    Pattern: "/",
    HandlerFunc: handlers.Index,
  },
  Route{
    Name: "TodoIndex",
    Method: "GET",
    Pattern: "/todos",
    HandlerFunc: handlers.TodoIndex,
  },
  Route{
    Name: "TodoShow",
    Method: "GET",
    Pattern: "/todos/{todoId:[0-9]}",
    HandlerFunc: handlers.TodoShow,
  },
  Route{
    Name: "TodoShow",
    Method: "PUT",
    Pattern: "/todos/{todoId:[0-9]}",
    HandlerFunc: handlers.CompleteTodo,
  },
  Route{
    Name: "TodoShow",
    Method: "DELETE",
    Pattern: "/todos/{todoId:[0-9]}",
    HandlerFunc: handlers.RemoveTodo,
  },
  Route{
    Name: "TodoShow",
    Method: "GET",
    Pattern: "/todos/active",
    HandlerFunc: handlers.ActiveTodo,
  },
  Route{
    Name: "TodoCreate",
    Method: "POST",
    Pattern: "/todos",
    HandlerFunc: handlers.TodoCreate,
  },
}

