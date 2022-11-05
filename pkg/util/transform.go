package util

import (
	"bufio"
	"io"

	"golang.org/x/text/encoding"
)

const transformer_buffer_size = 4096

func TransformFile(input io.Reader, output io.Writer, enc encoding.Encoding) error {
	transformer := enc.NewDecoder().Reader(input)

	out := bufio.NewWriter(output)

	// make a buffer to keep chunks that are read
	buf := make([]byte, transformer_buffer_size)

	total := 0

	for {
		// read a chunk
		n, err := transformer.Read(buf)

		total += n

		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}

		if _, err := out.Write(buf[:n]); err != nil {
			return err
		}
	}

	if err := out.Flush(); err != nil {
		return err
	}

	return nil
}
