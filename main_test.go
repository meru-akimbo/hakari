package main

import (
	"fmt"
	"testing"
)

func makeTestData() []map[string]string {
	return []map[string]string{
		{"tag": "test", "value": "1"},
		{"tag": "test", "value": "3"},
		{"tag": "test", "value": "-1"},
		{"tag": "test", "value": "-1.5"},
		{"hoge": "test", "value": "0"},
	}
}

func TestSum(t *testing.T) {
	result := sum("value", makeTestData())
	if result != 1.5 {
		t.Fatal("sum result is wrong: " + fmt.Sprint(result))
	}
}

func TestAve(t *testing.T) {
	result := ave("value", makeTestData())
	if result != 0.3 {
		t.Fatal("ave result is wrong: " + fmt.Sprint(result))
	}
}
