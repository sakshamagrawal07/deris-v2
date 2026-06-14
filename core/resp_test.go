package core

import (
	"fmt"
	"testing"
)

func TestSimpleStringDecode(t *testing.T) {
	cases := map[string]string{
		"+OK\r\n": "OK",
	}

	for k, v := range cases {
		value, _ := Decode([]byte(k))
		if value != v {
			t.Fail()
		}
	}
}

func TestErrorDecode(t *testing.T) {
	cases := map[string]string{
		"-ERR unknown command\r\n": "ERR unknown command",
	}

	for k, v := range cases {
		value, _ := Decode([]byte(k))
		if value != v {
			t.Fail()
		}
	}
}

func TestInt64Decode(t *testing.T) {
	cases := map[string]int64{
		":1000\r\n": 1000,
		":0\r\n":    0,
	}

	for k, v := range cases {
		value, _ := Decode([]byte(k))
		if value != v {
			t.Fail()
		}
	}
}

func TestBulkStringDecode(t *testing.T) {
	cases := map[string]string{
		"$6\r\nfoobar\r\n": "foobar",
		"$0\r\n\r\n":       "",
	}

	for k, v := range cases {
		value, _ := Decode([]byte(k))
		if value != v {
			t.Fail()
		}
	}
}

func TestArrayDecode(t *testing.T) {
	cases := map[string][]interface{}{
		"*0\r\n":                           {},
		"*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n": {"foo", "bar"},
		"*3\r\n:1\r\n:2\r\n:3\r\n":         {1, 2, 3},
		"*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$6\r\nfoobar\r\n":       {1, 2, 3, 4, "foobar"},
		"*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Foo\r\n-Bar\r\n": {[]int64{1, 2, 3}, []string{"Foo", "Bar"}},
	}

	for k, v := range cases {
		value, _ := Decode([]byte(k))
		array := value.([]interface{})
		if len(array) != len(v) {
			t.Fail()
		}
		for i := range array {
			if fmt.Sprintf("%v", v[i]) != fmt.Sprintf("%v", array[i]) {
				t.Fail()
			}
		}
	}
}
