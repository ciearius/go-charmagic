package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/cearius/go-charmagic"
	"github.com/cearius/go-charmagic/pkg/util"
	"github.com/spf13/cobra"
)

// transformCmd represents the guess command
var transformCmd = &cobra.Command{
	Use:   "transform",
	Short: "Convert a file to utf-8",
	Run:   TransformFile,
}

func ensureArgs(cmd *cobra.Command) (input, output string, err error) {
	inputFlag := cmd.Flag("input")
	outputFlag := cmd.Flag("output")

	if inputFlag.Value.String() == "" || outputFlag.Value.String() == "" {
		return "", "", errors.New("missing arguments")
	}

	output = outputFlag.Value.String()
	input = inputFlag.Value.String()

	return
}

func openFiles(input, output string) (inputFile, outputFile *os.File, err error) {
	outputFile, err = os.Create(output)

	if err != nil {
		return
	}

	inputFile, err = os.Open(input)

	if err != nil {
		return
	}

	return
}

func GetEncoding(cmd *cobra.Command) (string, error) {
	enc := cmd.Flag("encoding").Value.String()

	if enc != "" {
		return enc, nil
	}

	buf, err := util.ReadLinesAsBytes(cmd.Flag("input").Value.String())

	if err != nil {
		return "", err
	}

	result, err := charmagic.MatchBest(buf)

	if err != nil {
		return "", err
	}

	fmt.Printf("using %s with confidence of (%d/100)\n", result.Charset, result.Confidence)

	return result.Charset, nil
}

func TransformFile(cmd *cobra.Command, args []string) {
	input, output, err := ensureArgs(cmd)

	if err != nil {
		cmd.Usage()
		return
	}

	encName, err := GetEncoding(cmd)

	if err != nil {
		fmt.Printf("Failed to get encoding of input\n%s\n", err)
	}

	enc, err := charmagic.GetDecoder(encName)

	if err != nil {
		fmt.Printf("Failed to find a decoder for %s:\n%s\n", encName, err)
		return
	}

	inputFile, outputFile, err := openFiles(input, output)

	if err != nil {
		fmt.Printf("Failed to open output file.\n%s\n", err)
		return
	}

	defer inputFile.Close()
	defer outputFile.Close()

	err = util.TransformFile(inputFile, outputFile, enc)

	if err != nil {
		fmt.Printf("Failed to transform file.\n%s\n", err)
		return
	}

	fmt.Println("File written successfully")
}

func init() {
	rootCmd.AddCommand(transformCmd)

	transformCmd.Flags().StringP("encoding", "e", "", "source encoding")
	transformCmd.Flags().StringP("input", "i", "", "input file")
	transformCmd.Flags().StringP("output", "o", "", "output file")
}
