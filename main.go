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

		switch args[1] {
		case "sum":
			result = sum(args[2], records)

		case "ave":
			result = ave(args[2], records)
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

func ave(key string, ltsv []map[string]string) float64 {
	sum := sum(key, ltsv)
	return sum / float64(len(ltsv))
}
