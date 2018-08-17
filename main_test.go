package main

import (
	"testing"
)

func Test1Plus2Equals3(t *testing.T) {
	res := sum("1", "2")

	if res != "3" {
		t.Errorf("Sum was incorrect, got %s, want %s", res, "3")
	}
}

func Test1Plus1Equals2(t *testing.T) {
	res := sum("1", "1")

	if res != "2" {
		t.Errorf("Sum was incorrect, got %s, want %s", res, "2")
	}
}

func Test11Plus12Equals23(t *testing.T) {
	res := sum("11", "12")

	if res != "23" {
		t.Errorf("Sum was incorrect, got %s, want %s", res, "23")
	}
}

func Test11Plus19Equals30(t *testing.T) {
	res := sum("11", "19")

	if res != "30" {
		t.Errorf("Sum was incorrect, got %s, want %s", res, "30")
	}
}

func Test11Plus89Equals100(t *testing.T) {
	res := sum("11", "89")

	if res != "100" {
		t.Errorf("Sum was incorrect, got %s, want %s", res, "100")
	}
}

func Test11Plus189Equals200(t *testing.T) {
	res := sum("11", "189")

	if res != "200" {
		t.Errorf("Sum was incorrect, got %s, want %s", res, "200")
	}
}

func Test11Plus989Equals1000(t *testing.T) {
	res := sum("989", "11")

	if res != "1000" {
		t.Errorf("Sum was incorrect, got %s, want %s", res, "1000")
	}
}

func Test111Plus989Equals1100(t *testing.T) {
	res := sum("111", "989")

	if res != "1100" {
		t.Errorf("Sum was incorrect, got %s, want %s", res, "1100")
	}
}

func Test989Plus989Equals1978(t *testing.T) {
	res := sum("989", "989")

	if res != "1978" {
		t.Errorf("Sum was incorrect, got %s, want %s", res, "1978")
	}
}

func Test232Plus2Equals234(t *testing.T) {
	res := sum("232", "2")

	if res != "234" {
		t.Errorf("Sum was incorrect, got %s, want %s", res, "234")
	}
}
