/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var validStrings = [3]string{"snake_case", "cammelCase", "nospaces"}

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
	Run: func(cmd *cobra.Command, args []string) {
		// Can check with only flags...
		fmt.Println("normalize called")
		fmt.Println(args)
	},
	ValidArgs: validStrings[:],
	Args:      cobra.ExactValidArgs(1),
	Example:   "filename_normalizer normalize snake_case",
}

func init() {
	rootCmd.AddCommand(normalizeCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// normalizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// normalizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//normalizeCmd.Flags().StringP("type", "t", "", "Select the type of the arr")
	//normalizeCmd.MarkFlagRequired("type")
}
