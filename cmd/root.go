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
  version := "1.0.0"

  if _, err := os.Stat(homeDir+"/.todo-cache"); err == nil {
    byteCorbasi, _ := ioutil.ReadFile(homeDir+"/.todo-cache")
    listOfStructs := strings.Split(string(byteCorbasi), "/")

    for _, st := range listOfStructs {
      if st != "" {
        toDoList = append(toDoList, stringToStruct(st))
      }
    }
  }

  numberOfArguments := len(os.Args) - 1

  // exec command that have not args
  if numberOfArguments == 1 {
    if os.Args[1] == "-v" {
      fmt.Println(version)
    } else if os.Args[1] == "-h" {
      printHelp()
    } else if os.Args[1] == "-l" {
      currentList()
    } else if os.Args[1] == "-c" {
      completedList()
    } else {
      fmt.Println("Command you typed in doesn't exist")
      fmt.Println("Please enter todo -h for available commands.")
    }
  }

  // exec command that have Args
  if numberOfArguments == 2 {
    if os.Args[1] == "-a" {
      add()
    } else if os.Args[1] == "-m" {
      mark()
    } else if os.Args[1] == "-d" {
      deleteItem()
    } else {
      fmt.Println("Please enter todo -h for available commands.")
    }
  }

  // if it isn't 0, write a file to store toDoItem
  if len(toDoList) != 0 {
    ioutil.WriteFile(homeDir+"/.todo-cache", toByteConverter(toDoList), 0666)
  }
}
