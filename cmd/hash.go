/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Generate a hash of a string or file",
	Long: `Generate a hash of a string or file.
This command supports the following hash types:
- md5
- sha256
- sha512
You can provide the input string directly (--input) or read it from a file (--file).
The hash output can be printed to the console or saved to a file. (--output)
If --output is not provided, the generated hash will be printed to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		hashType, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Println("Error retrieving flag --type:", err)
			return
		}

		text, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Println("Error retrieving flag --input:", err)
			return
		}

		filePath, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println("Error retrieving flag --file:", err)
			return
		}

		savePath, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println("Error retrieving flag --output:", err)
			return
		}

		var inputText string

		if filePath != "" && text == "" {
			data, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			inputText = string(data)
		} else if text != "" && filePath == "" {
			inputText = text
		} else {
			fmt.Println("Please provide either an input string (--input) or a file path. (--file)")
			return
		}

		var hash string

		switch hashType {
		case "md5":
			hash = fmt.Sprintf("%x", md5.Sum([]byte(inputText)))
		case "sha256":
			hash = fmt.Sprintf("%x", sha256.Sum256([]byte(inputText)))
		case "sha512":
			hash = fmt.Sprintf("%x", sha512.Sum512([]byte(inputText)))
		default:
			fmt.Println("Please provide a valid hash type: <md5>, <sha256>, <sha512>")
			return
		}

		if savePath != "" {
			err := os.WriteFile(savePath, []byte(hash), 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
			fmt.Printf("Generated %s hash and saved to %s\n", hashType, savePath)
		} else {
			fmt.Println(hash)
		}
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)
	hashCmd.Flags().StringP("type", "t", "", "Hash type: <md5>, <sha256>, <sha512>")
	hashCmd.Flags().StringP("input", "i", "", "Input string to hash")
	hashCmd.Flags().StringP("file", "f", "", "File to read the input string from")
	hashCmd.Flags().StringP("output", "o", "", "Path to save the generated hash")
}
