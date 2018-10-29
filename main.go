package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"syscall"

	"github.com/najeira/ltsv"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	if terminal.IsTerminal(syscall.Stdin) {
		fmt.Println("Usage: cat hoge.ltsv | hakari (sum|ave) ${tareget_key}")
	} else {
		args := os.Args

		reader := ltsv.NewReader(os.Stdin)
		records, err := reader.ReadAll()

		if err != nil {
			panic(err)
		}

		var result float64

		switch args[1] {
		case "sum":
			result = sum(args[2], records)

		case "ave":
			result = ave(args[2], records)

		case "ile":
			per, _ := strconv.ParseFloat(args[3], 64)
			result = ile(args[2], per, records)
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
	return sum(key, ltsv) / float64(len(ltsv))
}

func ile(key string, per float64, ltsv []map[string]string) float64 {
	values := []float64{}

	for _, record := range ltsv {
		for k, v := range record {
			if k == key {
				f64, _ := strconv.ParseFloat(v, 64)
				values = append(values, f64)
			}
		}
	}

	sort.Float64s(values)
	length := float64(len(values))

	var num float64
	var few float64
	num = (length + 1.0) * per / 100.0
	few = num - float64(int(num))
	return values[int(num)-1.0] + few*(values[int(num)]-values[int(num)-1.0])
}
