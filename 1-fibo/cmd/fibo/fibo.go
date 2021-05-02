package main

import (
	"fmt"
	"github.com/phpclub/teachbase/1-fibo/pkg/fibolib"
)

var testNumbers = []int{1, 2, 3, 4, 5, 6, 100}

func main() {
	for _, v := range testNumbers {
		res, errStr := fibolib.FibByNum(v)
		if errStr != nil {
			fmt.Printf("%d => %s\n", v, errStr)
		} else {
			fmt.Printf("%d => %d\n", v, res)
		}
	}
}
