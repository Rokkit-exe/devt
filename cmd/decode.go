/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes a string using the specified encoding type",
	Long: `Decodes a string using the specified encoding type.
This command supports types:
- base64
- hex
You can provide the input string directly (--input) or read it from a file (--file).
The decoded output can be printed to the console or saved to a file. (--output)
If --output is not provided, the decoded string will be printed to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		decodeType, err := cmd.Flags().GetString("type")
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

		var decoded string
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
			fmt.Println("Please provide either --input or --file.")
			return
		}

		switch decodeType {
		case "base64":
			data, err := base64.StdEncoding.DecodeString(inputText)
			if err != nil {
				fmt.Println("Error decoding base64:", err)
				return
			}
			decoded = string(data)
		case "hex":
			data, err := hex.DecodeString(inputText)
			if err != nil {
				fmt.Println("Error decoding hex:", err)
				return
			}
			decoded = string(data)
		default:
			fmt.Println("Please provide a valid type: <base64|hex>")
			return
		}

		if savePath != "" {
			err := os.WriteFile(savePath, []byte(decoded), 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
			fmt.Printf("Decoded %s and saved to %s\n", decodeType, savePath)
		} else {
			fmt.Println(decoded)
		}
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	decodeCmd.Flags().StringP("type", "t", "", "Type of the encoding to decode")
	decodeCmd.Flags().StringP("input", "i", "", "Input string to decode")
	decodeCmd.Flags().StringP("file", "f", "", "File path to read the input string from")
	decodeCmd.Flags().StringP("output", "o", "", "File path to save the decoded output")
}
