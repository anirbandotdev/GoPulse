package cli

import (
	"fmt"

	"github.com/anirbandotdev/gopulse/constants"
	"github.com/anirbandotdev/gopulse/internal/audio"
	"github.com/spf13/cobra"
)

func init() {
	cmd := cobra.Command{
		Use:   "current",
		Short: "Get the current Profile",
		RunE: func(cmd *cobra.Command, _ []string) error {
			device, err := audio.GetBluetoothCard()
			if err != nil {
				return fmt.Errorf("Failed to detect Bluetooth device: %v", err)
			}
			profile, err := audio.GetCurrentProfile(device)
			if err != nil {
				return fmt.Errorf("Failed to get the current profile %v", err)
			}

			if profile == "" {
				fmt.Printf("No profile")
			}else if profile == string(constants.A2DP) {
				fmt.Printf("A2DP")
			}else {
				fmt.Printf("HFP/HSP")
			}
			
			return nil
		},
	}

	rootCmd.AddCommand(&cmd)
}
