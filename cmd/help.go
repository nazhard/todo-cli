package cmd

import "fmt"

func printHelp() {
  fmt.Println(`
------------------------------------------------------
         A Simple To Do
------------------------------------------------------
    -h          | list available commands
    -c          | list completed tasks
    -l          | list uncompleted tasks
    -v          | show version (it's good to chec          k your program)
    -a  <item>  | add <item> to list
    -m  <ID>    | mark item with <id> as complete
    -d  <ID>    | delete item with <id>
------------------------------------------------------
  `)
}
