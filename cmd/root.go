/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"strings"

	"aqw/internal/logic"
	"aqw/internal/window"

	"github.com/spf13/cobra"
)

var cfgName string // holds --config flag value

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aqwbot",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the window
		win, err := window.NewWindow()
		if err != nil {
			log.Fatalf("❌ Failed to create new window: %v", err)
		}

		// Choose config based on flag
		var cfg logic.FarmConfig
		switch strings.ToLower(cfgName) {
		case "", "default":
			cfg = logic.DefaultCfg
		case "cont":
			cfg = logic.ContinuousCfg
		case "warrior":
			cfg = logic.WarriorCfg
		case "healer":
			cfg = logic.HealerCfg
		default:
			log.Fatalf("❌ Unknown config %q. Valid options: default, warrior, healer", cfgName)
		}

		// Start farming
		if err := logic.ExecuteFarm(win, cfg); err != nil {
			log.Fatalf("❌ Farm encountered an error: %v", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// --config / -c selects which config to use
	rootCmd.PersistentFlags().StringVarP(
		&cfgName,
		"config",
		"c",
		"default",
		"Which config to use: default | warrior | healer",
	)

	// Example of another local flag (kept from your template)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
