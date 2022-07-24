/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var validTypes = [3]string{"snake_case", "cammelCase", "nospaces"}

// normalizeCmd represents the normalize command
var normalizeCmd = &cobra.Command{
	Use:   "normalize",
	Short: "Normalizes the name of the files and folder. ",
	Long: `Normalizes the name of the files and folders using a convention. The following conventions are supported: 
		- snake_case
		- cammelCase
		- nospaces
	To select the convention you want to use, use the --type= flag
	`,
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("normalize called")
		var directory string

		// Check for valid directory
		if len(args) == 1 {
			fileInfo, err := os.Stat(args[0])

			if err != nil || !fileInfo.IsDir() {
				return errors.New("argument given is not a valid directory")
			}

			directory = args[0]
			// Defaults to workging directory
		} else {
			directoryName, err := os.Getwd()

			if err != nil {
				return err
			}
			directory = directoryName
		}

		fmt.Printf("Using %s directory \n", directory)
		return nil
	},
	Args:    cobra.MaximumNArgs(1),
	Example: "filename_normalizer normalize -type=snake_case ./",
}

func init() {
	rootCmd.AddCommand(normalizeCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// normalizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// normalizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	normalizeCmd.Flags().StringP("type", "t", "", "Select the type of the arr")
	normalizeCmd.MarkFlagRequired("type")
}
