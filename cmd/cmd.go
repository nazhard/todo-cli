package cmd

import (
  "fmt"
  "io/ioutil"
  "os"
  "strconv"
  "strings"
  "time"
)

type toDoItem struct {
  op        string
  date      string
  completed string
}

func printHelp() {
  fmt.Println(`
-----------------------------------------------------
         A Simple To Do
-----------------------------------------------------
    -h		| list available commands
    -c		| list completed tasks
    -l		| list uncompleted tasks
    -v		| show version (it's good to check your program)
    -a  <item>	| add <item> to list
    -m  <ID>	| mark item with <id> as complete
    -d  <ID>	| delete item with <id>
-----------------------------------------------------
  `)
}

func toByteConverter(input []toDoItem) []byte {
  var stringified []string
  for _, td := range input {
    stringified = append(stringified, structToString(td))
  }
  return []byte(strings.Join(stringified, "/"))
}

func remove(slice []toDoItem, s int) []toDoItem {
  return append(slice[: s], slice[s+1:]...)
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

func Execute() {
  var toDoList []toDoItem
  var t toDoItem
  
  if _, err := os.Stat("cache"); err == nil {
    byteCorbasi, _ := ioutil.ReadFile("cache")
    listOfStructs := strings.Split(string(byteCorbasi), "/")
    for _, st := range listOfStructs {
      if st != "" {
        toDoList = append(toDoList, stringToStruct(st))
      }
    }
  }

  numberOfArguments := len(os.Args) - 1

  if numberOfArguments == 1 {
    if os.Args[1] == "-v" {
      fmt.Println("0.1.0")
    } else if os.Args[1] == "-h" {
      printHelp()
    } else if os.Args[1] == "-l" {
      fmt.Println("Task waiting to be done:\n")
      fmt.Println("   ID   | Item             | Date")
      fmt.Println("--------:------------------:----------")
      for i, td := range toDoList {
        if td.completed == "f" {
          fmt.Println("  ", i+1, "   |", td.op, "       |", td.date)
        }
      }
    } else if os.Args[1] == "-c" {
      fmt.Println("Your completed tasks:\n")
      fmt.Println("   ID   | Item               | Date")
      fmt.Println("--------:--------------------:----------")
      for i, td := range toDoList {
        if td.completed == "t" {
          fmt.Println("  ", i+1, "   |", td.op, "       |", td.date)
        }
      }
    } else {
      fmt.Println("Command you typed in does not exist.")
      fmt.Println("Please use ./todo-cli -h for available commands.")
    }
  } else if numberOfArguments == 2 {
    if os.Args[1] == "-a" {
      ti := time.Now()
      t.op = os.Args[2]
      t.completed = "f"
      t.date = ti.Format("02-01-2006")
      toDoList = append(toDoList, t)
    } else if os.Args[1] == "-m" {
      i, err := strconv.Atoi(os.Args[2])
      if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
      }
      i = i - 1
      toDoList[i].completed = "t"
    } else if os.Args[1] == "-d" {
      i, err := strconv.Atoi(os.Args[2])
      if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
      }
      i = i - 1
      /*if len(toDoList) == 1 {
				toDoList = toDoList[:0]
			} else {
			}*/
      toDoList = remove(toDoList, i)
      if len(toDoList) == 0 {
        ioutil.WriteFile("cache", toByteConverter(toDoList), 0666)
      }
    } else {
      fmt.Println("Command you typed in does not exist. Please use ./todo-cli -h for available commands.")
    }
  } else {
    fmt.Println("Please enter ./todo-cli -h for available commands.")
  }
  if len(toDoList) != 0 {
    ioutil.WriteFile("cache", toByteConverter(toDoList), 0666)
  }
}