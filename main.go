package main

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/sum", handleSum)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleSum(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query()["a"]
	b := r.URL.Query()["b"]

	log.Printf("a = %s", a)
	log.Printf("b = %s", b)
	sum := sum(strings.Join(a, ""), strings.Join(b, ""))
	log.Printf("sum = %s", sum)

	w.Write([]byte(sum))
}

func sum(a string, b string) string {
	if len(a) > len(b) {
		for i := 0; i <= int(math.Abs(float64(len(a)-len(b)))); i++ {
			b = "0" + b
		}
	} else if len(a) < len(b) {
		for i := 0; i <= int(math.Abs(float64(len(b)-len(a)))); i++ {
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
		ret = strconv.Itoa(carry) + ret[0:]
	}

	return ret
}
