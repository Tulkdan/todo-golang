package repos

import (
  "fmt"
  "github.com/tulkdan/todo/models"
)

var currentId int
var Todos models.Todos

func RepoFindTodo(id int) models.Todo {
  for _, t := range Todos {
    if t.Id == id {
      return t
    }
  }
  return models.Todo{}
}

func RepoCreateTodo(t models.Todo) models.Todo {
  currentId += 1
  t.Id = currentId
  Todos = append(Todos, t)
  return t
}

func RepoDestroyTodo(id int) error {
  for i, t := range Todos {
    if t.Id == id {
      Todos = append(Todos[:i], Todos[i+1:]...)
      return nil
    }
  }
  return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}

// Give us some seed data
func init() {
  RepoCreateTodo(models.Todo{Name: "Write presentation"})
  RepoCreateTodo(models.Todo{Name: "Host mettup"})
}
