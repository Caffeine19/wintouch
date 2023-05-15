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
			green.Println("create succeeded:\t", fileName)
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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wintouch",
	Short: "Use touch command like on unix",
	Long:  `Wintouch is used to bring the 'touch' command to windows powershell. As a command-line tool,it's built using Go`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(args)
		if len(args) > 0 {
			touchFile(args)
		} else {
			fmt.Println("😊This is wintouch, use -h for more help.")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wintouch.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
