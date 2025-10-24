package cli

import (
	"fmt"

	"github.com/anirbandotdev/gopulse/constants"
	"github.com/anirbandotdev/gopulse/internal/audio"
	"github.com/spf13/cobra"
)

func init() {
	cmd := cobra.Command{
		Use:   "toggle",
		Short: "Toggle Bluetooth audio profile between A2DP and HSP/HFP",
		RunE: func(cmd *cobra.Command, _ []string) error {
			device , err := audio.GetBluetoothCard()
			if err != nil {
				return fmt.Errorf("Failed to detect Bluetooth device: %v", err)
			}
			profile , err := audio.GetCurrentProfile(device)
			if err != nil{
				return fmt.Errorf("Failed to get the current profile %v" , err)
			}
			if profile == string(constants.A2DP) {
				if err := audio.SwitchProfile(device , string(constants.HSP)); err != nil{
					return fmt.Errorf("Failed to switch profile: %v", err)
				}
				fmt.Printf("✅ Switched from %s to %s\n", profile, string(constants.HSP))
			} else {
				if err := audio.SwitchProfile(device , string(constants.A2DP)); err!= nil {
					return fmt.Errorf("Failed to switch profile: %v", err)
				}
				fmt.Printf("✅ Switched from %s to %s\n", profile, string(constants.A2DP))
			}
			return nil
		},
	}

	rootCmd.AddCommand(&cmd)
}
