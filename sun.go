package main

////////////// IMPORTS //////////////////

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

///////// PACKAGE BUILD INFO ////////////

var (
    LocalInstall   string
    SysInstall     string
    PortableInstall string
	Version = "dev (version unknown)"
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
var statsjsonoutput bool
var readjsonoutput bool

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
	statsCmd.Flags().BoolVar(&statsjsonoutput, "json", false, "Format output in JSON")
	readCmd.Flags().BoolVar(&readjsonoutput, "json", false, "Format ouput in JSON")
}

//////////// COMMANDS VARS //////////////

// ROOT //
var rootCmd = &cobra.Command{
	Use:   "sun",
	Short: "Sun : a project information management system.",
	Long:  "Sun : a project information management system, built for coders. It manages langage, authors, mainteners, hosting and name of any project.",
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		GreetSun()
	},
}

// OTHERS //
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read .sunenv.yaml file.",
	Run: func(cmd *cobra.Command, args []string) {
		if !readjsonoutput {
			Read(false)
		} else {
			Read(true)
		}
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

var frameworkInitCmd = &cobra.Command{
	Use: "framework",
	Short: "Initialize frameworks config files.",
	Run: func(cmd *cobra.Command, args []string) {
		Ask()
	},
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Detect frameworks and langages.",
	Run: func(cmd *cobra.Command, args []string) {
		if statsjsonoutput {
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
	rootCmd.AddCommand(statsCmd)
	initCmd.AddCommand(frameworkInitCmd)
	if err := fang.Execute(context.TODO(), rootCmd); err != nil {
		os.Exit(1)
	}
}
