/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cearius/go-charmagic/pkg/api"
	"github.com/spf13/cobra"
)

// guessCmd represents the guess command
var guessCmd = &cobra.Command{
	Use:   "guess",
	Short: "Guess the encoding of a text-file",
	Run:   GuessEncoding,
}

func GuessEncoding(cmd *cobra.Command, args []string) {
	fileFlag := cmd.Flag("input")

	data, err := ReadLinesAsBytes(fileFlag.Value.String())

	if err != nil {
		fmt.Printf("Failed to read data from file.\n%s", err)
		return
	}

	res := api.MatchAll(data)

	for i, r := range res {
		fmt.Printf("%d. %d\t%s\t\t%s\n", i+1, r.Confidence, r.Charset, r.Language)
	}
}

const line_count = 10
const buf_size = 2048

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

		if line == line_count {
			break
		}

		if len(res) >= buf_size {
			break
		}
	}

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("encountered error while reading file: %s", err)
	}

	return res, nil
}

func init() {
	rootCmd.AddCommand(guessCmd)

	guessCmd.Flags().StringP("input", "i", "", "input file")
}
