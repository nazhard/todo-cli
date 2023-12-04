package cmd

import "fmt"

func completedList() {
  fmt.Println("Your completed tasks:")
  fmt.Println("   ID   | Item               | Date")
  fmt.Println("--------:--------------------:---------")
  for i, td := range toDoList {
    if td.completed == "t" {
      fmt.Println("  ", i+1, "   |", td.op, "|", td.date)
    }
  }
}

func currentList() {
  fmt.Println("Task waiting to be done:\n")
  fmt.Println("   ID   | Item             | Date")
  fmt.Println("--------:------------------:---------")
  for i, td := range toDoList {
    if td.completed == "f" {
      fmt.Println("  ", i+1, "   |", td.op, "       |", td.date)
    }
  }
}
