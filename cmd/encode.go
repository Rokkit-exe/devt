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

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		encodeType, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Println("Error retrieving flag <type>:", err)
			return
		}

		text, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Println("Error retrieving flag <text>:", err)
			return
		}

		filePath, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println("Error retrieving flag <file>:", err)
			return
		}

		savePath, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println("Error retrieving flag <save>:", err)
			return
		}

		var encoded string
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
			fmt.Println("Please provide either <text> or <file> to read the input string from.")
			return
		}

		switch encodeType {
		case "base64":
			encoded = base64.StdEncoding.EncodeToString([]byte(inputText))
		case "hex":
			encoded = hex.EncodeToString([]byte(inputText))
		default:
			fmt.Println("Please provide a valid type: <base64>, <hex>")
		}
		if savePath != "" {
			err := os.WriteFile(savePath, []byte(encoded), 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
			fmt.Printf("Encoded %s and saved to %s\n", encodeType, savePath)
		} else {
			fmt.Println(encoded)
		}
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	encodeCmd.Flags().StringP("type", "t", "", "Type of encoding (e.g., base64, hex)")
	encodeCmd.MarkFlagRequired("type")
	encodeCmd.Flags().StringP("input", "i", "", "Input string to encode")
	encodeCmd.Flags().StringP("file", "f", "", "Path to read the input string from")
	encodeCmd.Flags().StringP("output", "o", "", "Output file to save the encoded string")
}
