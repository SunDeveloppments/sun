package main

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Sun : a project information management system.",
	Long:  "Sun : a project information management system, built for coders. It manages langage, authors, mainteners, hosting and name of any project.",
	Run: func(cmd *cobra.Command, args []string) {
		GreetSun()
	},
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read .sunenv.yaml file.",
	Run: func(cmd *cobra.Command, args []string) {
		Read()
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize .sunenv.yaml.",
	Run: func(cmd *cobra.Command, args []string) {
		Init()
	},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show help.",
	Run: func(cmd *cobra.Command, args []string) {
		Help("main")
	},
}

var detectCmd = &cobra.Command{
	Use: "detect",
	Short: "Detect frameworks and langages.",
	Run: func(cmd *cobra.Command, args []string) {
		Detect()
		Frameworks()
	},
}

func main() {
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(detectCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
