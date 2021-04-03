package kvs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const help string = `Command	Usage\n
 -------	-----
 BEGIN		
 END
 `

func InitializeRepl() {
	reader := bufio.NewReader(os.Stdin)
	items := &Stack{}

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
			Set(operation[1], operation[2], items)
		case "GET":
			Get(operation[1], items)
		case "DELETE":
			Delete(operation[1], items)
		case "COUNT":
			Count(operation[1], items)
		case "QUIT":
			os.Exit(0)
		case "HELP":
			fmt.Printf("%s\n", help)
		default:
			fmt.Printf("Error: %s is an invalid command\n", operation[0])
		}
	}
}
