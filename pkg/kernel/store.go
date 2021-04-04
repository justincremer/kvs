package kernel

import "fmt"

type Dictionary map[string]interface{}

type Transaction struct {
	store Dictionary
	next  *Transaction
}

type Stack struct {
	top  *Transaction
	size int
}

type ITransaction interface {
	Push()
	Pop()
	Peek() *Transaction
	Commit()
	Rollback()
	Get(k string, T *Transaction)
	Set(k string, v string, T *Transaction)
	Delete(k string, s *Stack)
	Count(v string, s *Stack)
}

var Store = make(Dictionary)

func (stack *Stack) Push() {
	temp := Transaction{store: make(Dictionary)}
	temp.next = stack.top
	stack.top = &temp
	stack.size++
}

func (stack *Stack) Pop() {
	if stack.top != nil {
		node := &Transaction{}
		stack.top = stack.top.next
		node.next = nil
		stack.size--
		return
	}

	fmt.Printf("Error: No active transaction\n")
}

func (stack *Stack) Peek() *Transaction {
	return stack.top
}

func (stack *Stack) Commit() {
	if currTransaction := stack.Peek(); currTransaction != nil {
		for key, val := range currTransaction.store {
			Store[key] = val
			if currTransaction.next != nil {
				currTransaction.next.store[key] = val
			}
		}
		return
	}

	fmt.Printf("Nothing to commit\n")
}

func (stack *Stack) Rollback() {
	if stack.top != nil {
		for key := range stack.top.store {
			delete(stack.top.store, key)
		}
		return
	}

	fmt.Printf("Error: No active transaction\n")
}

func Get(key string, stack *Stack) {
	if currTransaction := stack.Peek(); currTransaction != nil {
		if val, ok := currTransaction.store[key]; ok {
			fmt.Printf("%s\n", val)
			return
		}

		fmt.Printf("Error: %s not set\n", key)
		return
	}

	if val, ok := Store[key]; ok {
		fmt.Printf("%s\n", val)
		return
	}

	fmt.Printf("Error: %s not set\n", key)
}

func Set(key string, val string, stack *Stack) {
	if currTransaction := stack.Peek(); currTransaction != nil {
		currTransaction.store[key] = val
		return
	}

	Store[key] = val
}

func Delete(key string, stack *Stack) {
	if currTransaction := stack.Peek(); currTransaction != nil {
		delete(currTransaction.store, key)
		fmt.Printf("%s deleted\n", key)
		return
	}

	delete(Store, key)
	fmt.Printf("%s deleted\n", key)
}

func Count(val string, stack *Stack) {
	var count int = 0
	if currTransaction := stack.Peek(); currTransaction != nil {
		for _, value := range currTransaction.store {
			if val == value {
				count++
			}
		}
		fmt.Printf("%d\n", count)
		return
	}

	for _, value := range Store {
		if val == value {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}
