package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/linuxsuren/jenkins-cli/app"
	"github.com/spf13/cobra"
)

type RootOptions struct {
	Version bool
}

var rootCmd = &cobra.Command{
	Use:   "jcli",
	Short: "jcli is a tool which could help you with your multiple Jenkins",
	Long: `jcli is Jenkins CLI which could help with your multiple Jenkins,
				  Manage your Jenkins and your pipelines
				  More information could found at https://jenkins-zh.cn`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Jenkins CLI (jcli) manage your Jenkins")

		current := getCurrentJenkins()
		if current != nil {
			fmt.Println("Current Jenkins is:", current.Name)
		} else {
			fmt.Println("Cannot found the configuration")
		}

		if rootOptions.Version {
			fmt.Printf("Version: v%.2f.%d%s", app.CurrentVersion.Number,
				app.CurrentVersion.PatchLevel,
				app.CurrentVersion.Suffix)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootOptions RootOptions

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&rootOptions.Version, "version", "v", false, "Print the version of Jenkins CLI")
}

func initConfig() {
	if err := loadDefaultConfig(); err != nil {
		if os.IsNotExist(err) {
			log.Printf("No config file found.")
			return
		}

		log.Fatalf("Config file is invalid: %v", err)
	}
}
