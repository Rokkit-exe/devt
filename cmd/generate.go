/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new random string or object",
	Long: `Generate a new random string or object.
This command supports the following types (--type):
- uuid
You can provide the output path to save the generated object. (--output)
If --output is not provided, the generated object will be printed to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		genType, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Println("Error retrieving flag <type>:", err)
			return
		}

		var text string
		switch genType {
		case "uuid":
			uuid := uuid.New()
			text = uuid.String()
		default:
			fmt.Println("Please provide a valid type: <uuid>")
		}

		savePath, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println("Error retrieving flag --output:", err)
			return
		}

		if savePath != "" {
			err := os.WriteFile(savePath, []byte(text), 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
			fmt.Printf("Generated %s and saved to %s\n", genType, savePath)
		} else {
			fmt.Println(text)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("type", "t", "", "Type of the object to generate")
	generateCmd.Flags().StringP("output", "o", "", "Output path to save the generated object")
	generateCmd.MarkFlagRequired("type")
}
