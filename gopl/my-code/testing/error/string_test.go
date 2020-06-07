package string

import (
	"testing"
)

func TestIsString(t *testing.T) {
	var trueTests = []struct {
		a    string
		want bool
		err  error
	}{
		{"abc", true, nil},
		{"xyx", true, nil},
	}

	var err error

	var falseTests = []struct {
		a    int
		want bool
		err  error
	}{
		{123, false, err},
		{123123, false, err},
	}

	for _, test := range trueTests {
		got, err := isString(test.a)
		if got != test.want || err != nil {
			t.Errorf("isString(%s) = (%v, %v), want (%v, %v)", test.a, got, err, test.want, test.err)
		}
	}

	for _, test := range falseTests {
		got, err := isString(test.a)
		if got != test.want || err == nil {
			t.Errorf("isString(%v) = (%v, %v), want (%v, %v)", test.a, got, err, test.want, test.err)
		}
	}

}
