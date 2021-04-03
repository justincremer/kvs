package kvs

import "fmt"

var (
	Store = make(map[string]string)
)

type Transaction struct {
	store map[string]string
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

func (s *Stack) Push() {
	temp := Transaction{store: make(map[string]string)}
	temp.next = s.top
	s.top = &temp
	s.size++
}

func (s *Stack) Pop() {
	if s.top != nil {
		node := &Transaction{}
		s.top = s.top.next
		node.next = nil
		s.size--
		return
	}

	fmt.Printf("Error: No active transaction\n")
}

func (s *Stack) Peek() *Transaction {
	return s.top
}

func (s *Stack) Commit() {
	if ct := s.Peek(); ct != nil {
		for k, v := range ct.store {
			Store[k] = v
			if ct.next != nil {
				ct.next.store[k] = v
			}
		}
		return
	}

	fmt.Printf("Nothing to commit\n")
}

func (s *Stack) Rollback() {
	if s.top != nil {
		for k := range s.top.store {
			delete(s.top.store, k)
		}
		return
	}

	fmt.Printf("Error: No active transaction\n")
}

func Get(k string, s *Stack) {
	if ct := s.Peek(); ct != nil {
		if v, ok := ct.store[k]; ok {
			fmt.Printf("%s\n", v)
			return
		}

		fmt.Printf("Error: %s not set\n", k)
		return
	}

	if v, ok := Store[k]; ok {
		fmt.Printf("%s\n", v)
		return
	}

	fmt.Printf("Error: %s not set\n", k)
}

func Set(k string, v string, s *Stack) {
	if ct := s.Peek(); ct != nil {
		ct.store[k] = v
		return
	}

	Store[k] = v
}

func Delete(k string, s *Stack) {
	if ct := s.Peek(); ct != nil {
		delete(ct.store, k)
		fmt.Printf("%s deleted\n", k)
		return
	}

	delete(Store, k)
	fmt.Printf("%s deleted\n", k)
}

func Count(v string, s *Stack) {
	var c int = 0
	if ct := s.Peek(); ct != nil {
		for _, val := range ct.store {
			if val == v {
				c++
			}
		}
		fmt.Printf("%d", c)
		return
	}

	for _, val := range Store {
		if val == v {
			c++
		}
	}

	fmt.Printf("%d", c)
}
