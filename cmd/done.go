package cmd

import (
  "fmt"
  "os"
  "strconv"
)

func done() {
  i, err := strconv.Atoi(os.Args[2])
  if err != nil {
    // handle error
    fmt.Println(err)
    os.Exit(2)
  }

  i = i - 1
  toDoList[i].completed = "t"
}
