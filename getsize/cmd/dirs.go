/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dirsCmd represents the dirs command
var dirsCmd = &cobra.Command{
	Use:   "dirs",
	Short: "Show the largest directories in the given path",
	Long:  "Show the largest directories in the given path",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dirs called")
	},
}

var Depth int
var Mindirsize int

func init() {
	rootCmd.AddCommand(dirsCmd)

	dirsCmd.PersistentFlags().IntVarP(&Depth, "depth", "", 2, "Depth of directory tree to display")
	viper.BindPFlag("depth", dirsCmd.PersistentFlags().Lookup("depth"))

	dirsCmd.PersistentFlags().IntVarP(&Mindirsize, "mindirsize", "", 100, "Only display directories larger than this threshold in MB.")
	viper.BindPFlag("mindirsize", dirsCmd.PersistentFlags().Lookup("mindirsize"))
}
