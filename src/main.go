package main

////////////// IMPORTS //////////////////

import (
	"os"

	"github.com/spf13/cobra"
)

//////// FLAGS DEFINITION SECTION ////////

var name string
var language string
var author string
var authorEmail string
var maintener string
var maintenerEmail string
var platform string
var repo string
var help bool
var y bool
var nohosting bool
var jsonoutput bool

func init() {
	initCmd.Flags().StringVar(&name, "name", "default", "The name of your package")
	initCmd.Flags().StringVar(&language, "language", "default", "The language in which your software is written")
	initCmd.Flags().StringVar(&author, "author", "default", "Your name")
	initCmd.Flags().StringVar(&authorEmail, "author-email", "default", "Email of author")
	initCmd.Flags().StringVar(&maintener, "maintener", "default", "Maintener of the repo")
	initCmd.Flags().StringVar(&maintenerEmail, "maintener-email", "default", "Email of maintener")
	initCmd.Flags().StringVar(&platform, "platform", "default", "Hosting platform")
	initCmd.Flags().StringVar(&repo, "repo", "default", "Repository URL")
	initCmd.Flags().BoolVar(&help, "help", false, "Show help")
	initCmd.Flags().BoolVar(&y, "y", false, "Confirm action without ask questions")
	initCmd.Flags().BoolVar(&nohosting, "no-hosting", false, "If the project has no hosting platform.")
	detectCmd.Flags().BoolVar(&jsonoutput, "json", false, "Format output in JSON")
}

//////////// COMMANDS VARS //////////////

var rootCmd = &cobra.Command{
	Use:   "sun",
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
		config := ConfigType{
			Name:           name,
			Language:       language,
			Author:         author,
			AuthorEmail:    authorEmail,
			Maintener:      maintener,
			MaintenerEmail: maintenerEmail,
			Platform:       platform,
			Repo:           repo,
		}
		Init(config, help, y, nohosting)
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
	Use:   "stats",
	Short: "Detect frameworks and langages.",
	Run: func(cmd *cobra.Command, args []string) {
		if jsonoutput {
			Detect(true)
		} else {
			Detect(false)
			Frameworks()
		}
	},
}

//////////////// MAIN //////////////////

func main() {
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(detectCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
