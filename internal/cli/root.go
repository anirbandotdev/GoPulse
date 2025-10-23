package cli

import (
	"log"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
    Use:   "gopulse",
    Short: "GoPulse — Simple Bluetooth Audio Profile Switcher",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
    }
}