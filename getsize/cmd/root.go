/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "getsize",
	Short: "List the size of the local directory",
	Long:  "This command will display the size of the local directory, with lots of options",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var Verbose bool
var Debug bool
var Highlight int
var Path string

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Debug output")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.PersistentFlags().IntVarP(&Highlight, "highlight", "", 500, "Hightlight files larger than threshold in MB")
	viper.BindPFlag("highlight", rootCmd.PersistentFlags().Lookup("highlight"))

	rootCmd.PersistentFlags().StringVarP(&Path, "path", "p", "", "Define the path to scan.")
	rootCmd.MarkPersistentFlagRequired("path")
	viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))
}
