package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/nwtgck/go-fakelish"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a username",
	Long: `Generate a username with a specific set of flagged options`,
	Run: generateUsername,
}

func generateUsername(cmd *cobra.Command, args []string) {
	minLengthArg, _ := cmd.Flags().GetInt("min")
	maxLengthArg, _ := cmd.Flags().GetInt("max")
	wordCountArg, _ := cmd.Flags().GetInt("words")
	separatorArg, _ := cmd.Flags().GetString("separator")
	includeNumbersArg, _ := cmd.Flags().GetBool("numbers")
	numbersCountArg, _ := cmd.Flags().GetInt("numbers-count")

	// Set the min and max length of the username with fallback
	minLength := minLengthArg | 6
	maxLength := maxLengthArg | 12
	wordCount := wordCountArg | 2
	separator := separatorArg

	if separator == "" {
		separator = "-"
	}

	// Initial string to append to
	initialString := ""

	// Loop 20 times
	for i := 0; i < wordCount; i++ {
		// Generate a fake word
		fakeWord := fakelish.GenerateFakeWord(minLength, maxLength)
		// Capitalize the first letter
		fakeWord = cases.Title(language.AmericanEnglish).String(fakeWord)
		// append the fake word
		initialString += fakeWord
	}

	// If the user wants numbers
	if includeNumbersArg {
		// Loop through the numbers count
		for i := 0; i < numbersCountArg; i++ {
			// Create ioReader
			ioReader := rand.Reader
			// Random integer between 0 and 9
			randomNumber, _ := rand.Int(ioReader, new (big.Int).SetInt64(9))
			// Append the number to the initial string
			initialString += fmt.Sprintf("%d", randomNumber)
		}
	}

	fmt.Println(initialString)
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
