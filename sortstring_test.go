package sortstring

import (
	"testing"
)

var results = []string{"132", "213", "231", "312", "321", "123"}

func TestReverse(t *testing.T) {
	a := NewSortString("racecar")
	a.Reverse(0, a.Len())
	if a.String() != "racecar" {
		t.Errorf("Failed palindrome test: %s is not %s\n", "racecar", a.String())
	}
}

func TestReverse2(t *testing.T) {
	a := NewSortString("abcdef")
	a.Reverse(0, a.Len())
	if a.String() != "fedcba" {
		t.Error("TestReverse2 failed")
	}
}

func TestBasicNextPermTest(t *testing.T) {
 	a := NewSortString("abcdef")
 	if !a.NextPermutation(0, a.Len()) {
 		t.Errorf("NextPermutation should have returned true! : %s\n", a)
 	}
}

func TestSortStringPerms(t *testing.T) {
	a := NewSortString("123")
	i := 0
	for a.NextPermutation(0, a.Len()) == true {
		t.Logf("Is %s == %s\n", a.String(), results[i])
		if results[i] != a.String() {
			t.FailNow()
		}
		i++
	}
	t.Logf("Is %s == %s\n", a.String(), results[i])
	if results[i] != a.String() {
		t.FailNow()
	}
}
