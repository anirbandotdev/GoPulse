package cli

import (
    "fmt"
    "github.com/anirbandotdev/gopulse/internal/audio"
    "github.com/spf13/cobra"
)

func init() {
    cmd := cobra.Command{
        Use:   "switch [profile]",
        Short: "Switch a Bluetooth device audio profile (use alias for profile : music/call/ad2p/hsp/hfp)",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
			profileInput := args[0]
			profile := audio.ResolveProfileAlias(profileInput)

			card, err := audio.GetBluetoothCard()
			if err != nil {
				return fmt.Errorf("Failed to detect Bluetooth device: %v", err)
			}
			if card == "" {
				return fmt.Errorf("No Bluetooth audio device found")
			}

			if err := audio.SwitchProfile(card, profile); err != nil {
				return fmt.Errorf("Failed to switch profile: %v", err)
			}

			fmt.Printf("✅ Switched %s to %s\n", card, profile)
			return nil
		},
    }
    rootCmd.AddCommand(&cmd)
	
}
