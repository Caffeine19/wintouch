/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func touchFile(fileNames []string) {
	for _, fileName := range fileNames {
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			_, err = os.Create(fileName)
			if err != nil {
				red := color.New(color.FgRed)
				red.Println("create failed:\t", err)
				os.Exit(1)
			}
			green := color.New(color.FgGreen)
			green.Printf("create succeeded:\t", fileName)
		} else {
			err = os.Chtimes(fileName, time.Now(), time.Now())
			if err != nil {
				red := color.New(color.FgRed)
				red.Println("update file date failed:\t", err)
				os.Exit(1)
			}
			blue := color.New(color.FgBlue)
			blue.Println("update file date succeeded:\t", fileName)
		}
	}
}

// touchCmd represents the touch command
var touchCmd = &cobra.Command{
	Use:   "touch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args", args)
		touchFile(args)
	},
}

func init() {
	rootCmd.AddCommand(touchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// touchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// touchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
