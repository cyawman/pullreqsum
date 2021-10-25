/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/cyawman/pullreqsum/internal/pullreqsum"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var repository string
var recipients string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pullreqsum",
	Short: "Github Pull Request Summary Fetcher",
	Long:  `Fetch the Github pull requests created for a given repository and email the results to one or more recipients`,
	Run: func(cmd *cobra.Command, args []string) {

		config := pullreqsum.Config{
			GithubRepository:  viper.GetString("github.repository"),
			MessageSender:     viper.GetString("mailer.message.sender"),
			MessageRecipients: viper.GetStringSlice("mailer.message.recipients"),
			MessageSubject:    viper.GetString("mailer.message.subject"),
		}

		pullreqsum.Run(config)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pullreqsum.yaml)")
	rootCmd.PersistentFlags().StringVar(&repository, "repository", "", "Github.com repository {{owner}}/{{repo}} (Example: golang/go)")
	rootCmd.PersistentFlags().StringVar(&recipients, "recipients", "", "List of email addresses to recieve the summary report")

	viper.BindPFlag("github.repository", rootCmd.PersistentFlags().Lookup("repository"))
	viper.BindPFlag("mailer.message.recipients", rootCmd.PersistentFlags().Lookup("recipients"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".pullreqsum" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("pullreqsum")
	}

	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
