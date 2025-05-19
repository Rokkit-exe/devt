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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
			fmt.Println("Error retrieving flag <save>:", err)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().StringP("type", "t", "", "Type of the object to generate")
	generateCmd.Flags().StringP("output", "o", "", "Output path to save the generated object")
	generateCmd.MarkFlagRequired("type")
}
