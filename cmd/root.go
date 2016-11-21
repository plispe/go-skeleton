package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-skelleton",
	Short: "Go appliaction skeleton",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// func init() {
// 	cobra.OnInitialize(initConfig)
//
// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports Persistent Flags, which, if defined here,
// 	// will be global for your application.
//
// 	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-skelleton.yaml)")
// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" { // enable ability to specify config file via flag
// 		viper.SetConfigFile(cfgFile)
// 	}
//
// 	viper.SetConfigName(".go-skelleton") // name of config file (without extension)
// 	viper.AddConfigPath("$HOME")         // adding home directory as first search path
// 	viper.AutomaticEnv()                 // read in environment variables that match
//
// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
