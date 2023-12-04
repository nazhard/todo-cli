package cmd

import (
  "fmt"
  "io/ioutil"
  "os"
  "strconv"
)

func remove(slice []toDoItem, s int) []toDoItem {
  return append(slice[: s], slice[s+1:]...)
}

func deleteItem() {
  i, err := strconv.Atoi(os.Args[2])
  if err != nil {
    // handle error
    fmt.Println(err)
    os.Exit(2)
  }

  i = i - 1

  toDoList = remove(toDoList, i)
  if len(toDoList) == 0 {
    ioutil.WriteFile("cache", toByteConverter(toDoList), 0666)
  }
}
