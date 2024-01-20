package cmd

import "fmt"

func printHelp() {
  fmt.Println(`
  Simple To-Do

  Usage:
    todo [options] [args]

  Options:
    h, help
      print this help message
    c, complete
      list completed tasks
    l, list
      list uncompleted tasks
    v, version
      show version
    a <item>, add <item>
      add <item> to list
    d <id>, done <id>
      mark item with <id> as completed
    r <id>, remove <id>
      remove item with <id>
      
  Example:
    todo v
    todo add "simple task"
`)
}
