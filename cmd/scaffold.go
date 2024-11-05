/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ylanzinhoy/internal/controller"
)

// scaffoldCmd represents the scaffold command

var packageFlag string

var scaffoldCmd = &cobra.Command{
	Use:   "scaffold [entity/model Name] [fields:type] [package name]",
	Short: "A brief description of your command",
	Args:  cobra.MinimumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		controller.ScaffoldController(args, packageFlag)

	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)

	scaffoldCmd.Flags().StringVar(&packageFlag, "package", "com.example.demo", "Nome do pacote para a classe Java")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scaffoldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scaffoldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
