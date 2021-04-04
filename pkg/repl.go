package kvs

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/justincremer/kvs/pkg/kernel"
)

const help string = `Command    Usage
-------    -----
BEGIN      Pushes a new transaction stream onto the stack
END        Pops the current transaction stream off of the stack
COMMIT     Commits current transaction head and sub transactions to the global store
ROLLBACK   Rolls back the last commit
GET        Returns a key value from either the current transaction or global store
SET        Sets a key value in the current transaction
DELETE     Deletes a key value from either the current transaction or global store
COUNT      Returns the stack size of the current transaction if there is one
           If there isn't one, it will return the stack size of the global store			
HELP       If you're using this program, you know what this does
QUIT       Exit 0`

func InitializeRepl() {
	reader := bufio.NewReader(os.Stdin)
	items := &kernel.Stack{}

	for {
		fmt.Printf("$ ")
		text, _ := reader.ReadString('\n')
		operation := strings.Fields(text)

		switch operation[0] {
		case "BEGIN":
			items.Push()
		case "END":
			items.Pop()
		case "COMMIT":
			items.Commit()
		case "ROLLBACK":
			items.Rollback()
		case "SET":
			kernel.Set(operation[1], operation[2], items)
		case "GET":
			kernel.Get(operation[1], items)
		case "DELETE":
			kernel.Delete(operation[1], items)
		case "COUNT":
			kernel.Count(operation[1], items)
		case "SAVE":
			kernel.Save("daddy.db")
		case "LOAD":
			kernel.Load("daddy.db")
		case "QUIT":
			os.Exit(0)
		case "HELP":
			fmt.Printf("%s\n", help)
		default:
			fmt.Printf("Error: %s is an invalid command\n", operation[0])
		}
	}
}
