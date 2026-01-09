package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"{{ cookiecutter.module_name }}/logger"
	"{{ cookiecutter.module_name }}/version"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "{{ cookiecutter.bin_name }}",
	Short:   "A brief description of your application",
	Long:    `A longer description`,
	// Version: fmt.Sprintf("%s", version.GetVersion()),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Init(viper.GetString("log.level"), viper.GetBool("log.caller"), viper.GetString("log.file"), viper.GetBool("log.json"))
	},
{% if cookiecutter.subcommands != "y" %}
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
{% endif %}
}

// subcommands version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version details",
	Long:  `Print the version details including build date, git ref, SHA256, OS and Arch.`,
	Run: func(cmd *cobra.Command, args []string) {
		v := version.GetVersion()
		fmt.Print(v.Details())
	},
}

// subcommands logtest
var logtestCmd = &cobra.Command{
	Use:   "logtest",
	Short: "Test logging functionality with different log levels",
	Long:  `This command demonstrates logging with different levels (debug, info, warn, error, fatal).`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("This is a debug message - only shown when log level is debug")
		log.Info().Msg("This is an info message - shown when log level is info or lower")
		log.Warn().Msg("This is a warning message - shown when log level is warn or lower")
		log.Error().Msg("This is an error message - shown when log level is error or lower")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Config
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "{{ cookiecutter.config_type }} config file")

	// Logging - make all log flags persistent so they're available to subcommands
	rootCmd.PersistentFlags().String("log-level", "info", "Set the log level (debug, info, warn, error, fatal, panic)")
	viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))

	rootCmd.PersistentFlags().String("log-file", "", "Write logs in json format to this file")
	viper.BindPFlag("log.file", rootCmd.PersistentFlags().Lookup("log-file"))

	rootCmd.PersistentFlags().Bool("log-caller", false, "Include the caller file and line number")
	viper.BindPFlag("log.caller", rootCmd.PersistentFlags().Lookup("log-caller"))

	rootCmd.PersistentFlags().Bool("log-json", false, "Log as json messages")
	viper.BindPFlag("log.json", rootCmd.PersistentFlags().Lookup("log-json"))

	// subcommands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(logtestCmd)
}
