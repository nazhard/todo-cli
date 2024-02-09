package cmd

import (
  "os"
  "strconv"
)

func done() {
  i := os.Args[2]

  h, _ := strconv.Atoi(i)

  h = h - 1
  toDoList[h].completed = "t"
}
