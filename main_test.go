package main

import (
	"fmt"
	"strconv"
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

func makeIleTestData() []map[string]string {
	return []map[string]string{
		{"tag": "test", "value": "16"},
		{"tag": "test", "value": "19"},
		{"tag": "test", "value": "24"},
		{"tag": "test", "value": "31"},
		{"tag": "test", "value": "40"},
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

func TestIle(t *testing.T) {
	result_str := fmt.Sprintf("%.5f", ile("value", 60.0, makeIleTestData()))
	result, _ := strconv.ParseFloat(result_str, 64)

	if result != 28.2 {
		t.Fatal("ile result is wrong: " + fmt.Sprint(result))
	}
}
