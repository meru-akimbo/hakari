package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/najeira/ltsv"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	if terminal.IsTerminal(1) {
		args := os.Args

		reader := ltsv.NewReader(os.Stdin)
		records, _ := reader.ReadAll()

		var result float64
		if args[1] == "sum" {
			result = sum(args[2], records)
		}

		fmt.Println(result)
	}
}

func sum(key string, ltsv []map[string]string) float64 {
	result := 0.0
	for _, record := range ltsv {
		for k, v := range record {
			if k == key {
				f64, _ := strconv.ParseFloat(v, 64)
				result += f64
			}
		}
	}

	return result
}
