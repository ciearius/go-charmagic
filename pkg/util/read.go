package util

import (
	"bufio"
	"fmt"
	"os"
)

const line_count = 10
const buf_size = 4096

func ReadLinesAsBytes(file string) ([]byte, error) {
	res := []byte{}
	buf := make([]byte, buf_size)
	line := 0

	input, err := os.Open(file)

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s", err)
	}

	sc := bufio.NewScanner(input)

	sc.Split(bufio.ScanLines)
	sc.Buffer(buf, buf_size)

	for {
		if ok := sc.Scan(); !ok {
			break
		}

		res = append(res, sc.Bytes()...)
		line++

		if line == line_count || len(res) >= buf_size {
			break
		}
	}

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("encountered error while reading file: %s", err)
	}

	return res, nil
}
