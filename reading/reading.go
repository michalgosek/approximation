package reading

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func FromInput() (float64, error) {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("type degree: ")
	for scanner.Scan() {
		if scanner.Text() != "0" {
			input = scanner.Text()
			break
		}
	}

	x, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.WithMessage(err, "parsing to float64 failed")
	}
	return x, nil
}

func FileToMap(path string) (map[int][]float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithMessage(err, "cannot open file")
	}

	fmt.Println("data.csv has been sucessfuly read")
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return nil, errors.WithMessage(err, "read all failed")
	}

	m := make(map[int][]float64, len(records))
	for i, numbers := range records {

		var values []float64
		for _, n := range numbers {
			v, err := strconv.ParseFloat(n, 32)
			if err != nil {
				return nil, errors.WithMessage(err, "parsing to float64 failed")
			}

			values = append(values, v)
		}

		m[i+1] = values
	}

	fmt.Println("data has been loaded into map")
	return m, nil
}
