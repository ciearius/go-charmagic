package cmd

import (
	"fmt"

	"github.com/cearius/go-charmagic"
	"github.com/cearius/go-charmagic/pkg/util"
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

	data, err := util.ReadLinesAsBytes(fileFlag.Value.String())

	if err != nil {
		fmt.Printf("Failed to read data from file.\n%s", err)
		return
	}

	res := charmagic.MatchAll(data)

	for i, r := range res {
		fmt.Printf("%d. %d\t%s\t\t%s\n", i+1, r.Confidence, r.Charset, r.Language)
	}
}

func init() {
	rootCmd.AddCommand(guessCmd)
	guessCmd.Flags().StringP("input", "i", "", "input file")
}
