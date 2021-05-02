package fibolib

import (
	"errors"
	"fmt"
)

const MaxNum = 20
const NumberNotValid = "ошибка: максимальный номер числа Фибоначи не должен превышать"

// FibByNum вычисляем число Фибоначи по номеру
func FibByNum(n int) (int, error) {
	var a, b = 1, 1
	if n > MaxNum {
		return 0, errors.New(fmt.Sprintf("%s : %d", NumberNotValid, MaxNum))
	}

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a, nil
}
