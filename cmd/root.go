package cmd

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

type toDoItem struct {
  op        string
  date      string
  completed string
}

var (
  toDoList []toDoItem
  t toDoItem
  homeDir, _ = os.UserHomeDir()
)

func toByteConverter(input []toDoItem) []byte {
  var stringified []string
  for _, td := range input {
    stringified = append(stringified, structToString(td))
  }

  return []byte(strings.Join(stringified, "/"))
}

func structToString(input toDoItem) string {
  if input.op != "" {
      str := fmt.Sprintf("%s", input.op) + "|" + fmt.Sprintf("%s", input.date) + "|" + fmt.Sprintf("%s", input.completed)

    return str
  } else {
    return ""
  }
}

func stringToStruct(input string) toDoItem {
  inp := strings.Split(input, "|")

  var t toDoItem

  if len(inp) == 3 {
    t.op = inp[0]
    t.date = inp[1]
    t.completed = inp[2]
  }

  return t
}

// exec all cmd
func Execute() {
  hello := "Welcome to a simple To-Do CLI. Dev by Nazhard. \ntry type todo help"
  version := "1.0.0"

  if _, err := os.Stat(homeDir+"/.todo"); err == nil {
    byteCorbasi, _ := ioutil.ReadFile(homeDir+"/.todo")
    listOfStructs := strings.Split(string(byteCorbasi), "/")

    for _, st := range listOfStructs {
      if st != "" {
        toDoList = append(toDoList, stringToStruct(st))
      }
    }
  }

  numberOfArguments := len(os.Args) - 1

  if len(os.Args) == 1 {
    fmt.Println(hello)
  }

  // exec command that have not args
  if numberOfArguments == 1 {
    if os.Args[1] == "v" || os.Args[1] == "version" {
      fmt.Println(version)
    } else if os.Args[1] == "h" || os.Args[1] == "help"{
      printHelp()
    } else if os.Args[1] == "l" || os.Args[1] == "list" {
      currentList()
    } else if os.Args[1] == "c" || os.Args[1] == "complete" {
      completedList()
    } else {
      fmt.Println("Command you typed in doesn't exist")
      fmt.Println("Please enter \"todo help\" for available commands.")
    }
  }

  // exec command that have Args
  if numberOfArguments == 2 {
    if os.Args[1] == "a" || os.Args[1] == "add"{
      add()
    } else if os.Args[1] == "d" || os.Args[1] == "done" {
      done()
    } else if os.Args[1] == "r" || os.Args[1] == "remove" {
      removeItem()
    } else {
      fmt.Println("Please enter \"todo help\" for available commands.")
    }
  }

  // if it isn't 0, write a file to store toDoItem
  if len(toDoList) != 0 {
    ioutil.WriteFile(homeDir+"/.todo", toByteConverter(toDoList), 0666)
  }
}
