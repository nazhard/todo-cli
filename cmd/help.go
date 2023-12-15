package cmd

import "fmt"

func printHelp() {
  fmt.Println(`
------------------------------------------------------
         A Simple To Do
------------------------------------------------------

[USAGE]

  todo [COMMANDS] | [ITEM]

[COMMANDS]

  -h | help               | print this help message
  -c | complete           | list completed tasks
  -l | list               | list uncompleted tasks
  -v | version            | show version
  -a <item> | add <item>  | add <item> to list
  -d <id>   | done <id>   | mark item with <id> as completed
  -r <id>   | remove <id> | remove item with <id>

[EXAMPLE]

  todo -v | todo add "simple task"
  `)
}
