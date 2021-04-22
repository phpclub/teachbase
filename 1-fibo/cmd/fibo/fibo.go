package main

import (
	"fmt"
	"os"
)

const maxNum = 20

func main() {
	var res, errStr = FibByNum(10)
	if errStr != nil {
		fmt.Println(errStr)
		os.Exit(1)
	}
	fmt.Println(res)

}
