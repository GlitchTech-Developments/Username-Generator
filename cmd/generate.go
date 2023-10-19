/*
Copyright © 2023 GlitchTech Developments <dev@glitchtech.eu>
*/
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
	minLength, _ := cmd.Flags().GetInt("min")
	maxLength, _ := cmd.Flags().GetInt("max")
	wordCount, _ := cmd.Flags().GetInt("words")
	separator, _ := cmd.Flags().GetString("separator")
	includeNumbers, _ := cmd.Flags().GetBool("numbers")

	// Initial string to append to
	initialString := ""
	for i := 0; i < wordCount; i++ {
		// Generate a fake word
		fakeWord := fakelish.GenerateFakeWord(minLength, maxLength)
		// Capitalize the first letter
		fakeWord = cases.Title(language.AmericanEnglish).String(fakeWord)
		
		// If it's not the first or last word
		if i != 0 {
			
			// Append the separator
			initialString += separator
			// Append the fake word
			initialString += fakeWord
		} else {
			initialString += fakeWord
		}
	}

	// If the user wants numbers
	if includeNumbers {
		// Loop through the numbers count
		for i := 0; i < 4; i++ {
			// Create ioReader
			ioReader := rand.Reader
			// Random integer between 0 and 9
			randomNumber, _ := rand.Int(ioReader, new (big.Int).SetInt64(9))
			// Append the number to the initial string
			initialString += fmt.Sprintf("%d", randomNumber)
		}
	}

	fmt.Println("Done! Your generated username is: ", initialString)
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("min", "m", 6, "Minimum length of the username")
	generateCmd.Flags().IntP("max", "M", 12, "Maximum length of the username")
	generateCmd.Flags().IntP("words", "w", 2, "Number of words to generate")
	generateCmd.Flags().StringP("separator", "s", "", "Separator between words")
	generateCmd.Flags().BoolP("numbers", "n", false, "Include numbers in the username")
}
