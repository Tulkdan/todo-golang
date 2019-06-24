package handlers

import (
  "io"
  "encoding/json"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/tulkdan/todo/repos"
  "github.com/tulkdan/todo/models"
  "strconv"
)

func ActiveTodo(w http.ResponseWriter, r *http.Request) {
  var todos = []models.Todo{}
  for _, t := range repos.Todos {
    if t.Completed {
      todos = append(todos, t)
    }
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  json.NewEncoder(w).Encode(todos)
}

func CompleteTodo(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoId, _ := strconv.Atoi(vars["todoId"])
  todo := models.Todo{}

  if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  for index, t := range repos.Todos {
    if t.Id == todoId {
      t.Completed = todo.Completed
      repos.Todos = append(repos.Todos[:index], repos.Todos[index + 1:]...)
      repos.Todos = append(repos.Todos, t)
      todo = t
      break
    }
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  if todo.Id == 0 && todo.Name == "" {
    w.WriteHeader(404)
  }
  json.NewEncoder(w).Encode(todo)
}


func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome!")
}

func RemoveTodo(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoId, _ := strconv.Atoi(vars["todoId"])
  for index, t := range repos.Todos {
    if t.Id == todoId {
      repos.Todos = append(repos.Todos[:index], repos.Todos[index + 1:]...)
      break
    }
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  io.WriteString(w, `{"message": "Todo removed!"}`)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
  var todo models.Todo
  if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  if err := r.Body.Close(); err != nil {
    panic(err)
  }

  t := repos.RepoCreateTodo(todo)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  if err := json.NewEncoder(w).Encode(t); err != nil {
    panic(err)
  }
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(repos.Todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoId, _ := strconv.Atoi(vars["todoId"])
  todo := models.Todo{}
  for _, t := range repos.Todos {
    if t.Id == todoId {
      todo = t
      break
    }
  }
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  json.NewEncoder(w).Encode(todo)
}
