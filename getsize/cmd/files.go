/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Show the largest file in the given path",
	Long:  "Quickly find the largest file in the given path",

	Run: func(cmd *cobra.Command, args []string) {
		for key, value := range viper.GetViper().AllSettings() {
			log.WithFields(log.Fields{
				key: value,
			}).Info("command flag")
		}
	},
}

var Filecount int
var Minfilesize int64

func init() {
	rootCmd.AddCommand(filesCmd)

	filesCmd.PersistentFlags().IntVarP(&Filecount, "filecount", "f", 10, "Limit the number of files to display")
	viper.BindPFlag("filecount", filesCmd.PersistentFlags().Lookup("filecount"))

	filesCmd.PersistentFlags().Int64VarP(&Minfilesize, "minfilesize", "", 50, "Minimum size for files in search in MB")
	viper.BindPFlag("minfilesize", filesCmd.PersistentFlags().Lookup("minfilesize"))
}
