/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var NumberOfTimes int

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo STRING",
	Short: "Echo string which is passed as an argument",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := echo(cmd, args[0]); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
	echoCmd.Flags().IntVarP(&NumberOfTimes, "number", "n", 1, "Number of times to echo ssss")
}

func echo(cmd *cobra.Command, str string) error {
	for i := 0; i < NumberOfTimes; i++ {
		fmt.Println(str)
	}
	return nil
}
