package cmd

import (
  "os"
  "time"
)

func add() {
  ti := time.Now()
  t.op = os.Args[2]
  t.completed = "f"
  t.date = ti.Format("02-01-2006")
  toDoList = append(toDoList, t)
}
