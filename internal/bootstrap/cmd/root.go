package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	appBuild, appVersion string

	rootCmd = &cobra.Command{
		Use:               "grafo",
		Short:             "GrafoDB is a fast distributed graph database",
		PersistentPreRunE: initialize,
	}
)

func Execute(build, version string) error {
	appBuild = build
	appVersion = version

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().BoolP(flagDebug, flagDebugShort, false, flagDebugDesc)
	rootCmd.PersistentFlags().StringP(flagLogFormat, flagLogFormatShort, defaultLogFormat, flagLogFormatDesc)
	rootCmd.PersistentFlags().Bool(flagVerbose, false, flagVerboseDesc)
}

func initialize(cmd *cobra.Command, _ []string) error {
	// cli
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	envKeyReplacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(envKeyReplacer)

	viper.SetEnvPrefix("GRAFO")
	viper.AutomaticEnv()

	// logging
	if viper.GetBool(flagVerbose) {
		viper.Set(flagDebug, true)
	}

	if viper.GetBool(flagDebug) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	logFormat := viper.GetString(flagLogFormat)
	switch logFormat {
	case logFormatConsole:
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	case logFormatJSON:
		zerolog.TimeFieldFormat = ""
	default:
		return fmt.Errorf("invalid log format: %s", logFormat)
	}

	log.Logger = log.Logger.With().
		Str("app", "grafodb").
		Str("version", appVersion).
		Str("build", appBuild).
		Logger()
	rootLogger = &log.Logger

	return nil
}
