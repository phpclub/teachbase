package fibolib

import "errors"

const MaxNum = 20

// FibByNum вычисляем число Фибоначи по номеру
func FibByNum(n int) (int, error) {
	var a, b = 1, 1
	if n > MaxNum {
		return 0, errors.New("ошибка: максимальный номер числа Фибоначи не должен превышать 20")
	}

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a, nil
}
