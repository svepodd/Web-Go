package main

import (
	"storage/list"
)

func main() {
	l := list.List{}
	l.Add(12)
	l.Add(65)
	l.Add(99)
	l.Add(99)
	l.Add(787)
	l.Add(12)
	l.Add(12)
	l.Add(12)
	l.Print()
}
