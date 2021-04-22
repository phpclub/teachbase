package main

import (
	"fmt"
	"github.com/phpclub/teachbase/1-fibo/pkg/fibolib"
	"os"
)

func main() {
	var res, errStr = fibolib.FibByNum(10)
	if errStr != nil {
		fmt.Println(errStr)
		os.Exit(1)
	}
	fmt.Println(res)

}
