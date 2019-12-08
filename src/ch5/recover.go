package main

import "fmt"

func Recover() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}

	}()
	foo()
	return err
}

func foo() {
	panic("application panic!!!")
}
