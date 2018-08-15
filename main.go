package main

import (
	"math"
	"strconv"
)

func main() {

}

func sum(a string, b string) string {
	if len(a) > len(b) {
		for i := 0; i < int(math.Abs(float64(len(a)-len(b)))); i++ {
			b = "0" + b
		}
	} else if len(a) < len(b) {
		for i := 0; i < int(math.Abs(float64(len(a)-len(b)))); i++ {
			a = "0" + a
		}
	}
	pos := len(a) - 1
	ret := ""
	carry := 0
	var digit int
	for pos >= 0 {
		digit = int(a[pos]-'0') + int(b[pos]-'0')
		if carry > 0 {
			digit += carry
			carry = 0
		}
		if digit >= 10 {
			carry++
			digit = digit % 10
		}
		ret = strconv.Itoa(digit) + ret
		pos--
	}

	if carry > 0 {
		// if int(ret[0]-'0') != 0 {
		// 	digit = carry + int(ret[0]-'0')
		// 	panic(ret)
		// 	ret = strconv.Itoa(digit) + ret[1:]
		// } else {
		// 	ret = strconv.Itoa(carry) + ret[0:]
		// }
		ret = strconv.Itoa(carry) + ret[0:]
	}

	return ret
}
