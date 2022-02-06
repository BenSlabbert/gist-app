package cmd

import (
	"fmt"
	"github.com/BenSlabbert/gist-app/pkg/env"
	"github.com/BenSlabbert/gist-app/pkg/githubapi"
	"github.com/spf13/cobra"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var api *githubapi.Api

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gist-app",
	Short: "Manage your Github Gists from the terminal!",
	Long: `
This cli allows you to perform simple CRUD on your Github Gists.
You need to provide Github Username as well as a Github Access token, create tokens here: https://github.com/settings/tokens/new.
For your peace of mind, limit the access of the token to only Gists.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gist-app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var err error

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		var home string
		home, err = homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gist-app" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gist-app")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err = viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	cobra.CheckErr(err)

	log.Printf("username from config: %s", viper.Get(env.GithubApiUsername))
	api, err = githubapi.NewApi(viper.GetString(env.GithubApiUsername), viper.GetString(env.GithubApiToken))
	cobra.CheckErr(err)
}
