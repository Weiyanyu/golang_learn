package main

type IntList struct {
	Value int
	Next  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Next.Sum()
}
