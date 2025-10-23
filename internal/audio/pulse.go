package audio

import (
	"os/exec"
	"strings"
)

func GetSinks() ([]string, error) {

	cmd := exec.Command("pactl", "list", "sinks", "short") // pactl list cards short (alternative way)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	var sinks []string

	for _, line := range lines {
		if line != "" {
			fields := strings.Fields(line)
			if len(fields) > 1 {
				sinks = append(sinks, fields[1])
			}
		}
	}
	return sinks, nil
}

func SwitchProfile(device, profile string) error {
	cmd := exec.Command("pactl", "set-card-profile", device, profile)
	return cmd.Run()
}

func GetCurrentProfile(device string) (string , error) {
	cmd := exec.Command("pactl" , "list" , "cards")
	out , err := cmd.Output()
	if err != nil {
        return "", err
    }

	lines := strings.Split(string(out) , "\n")
	var current string
	
	for _ , line := range lines{
		if strings.Contains(line , device) {
			current = device
		}

		if current != "" && strings.Contains(line , "Active Profile: "){
			return strings.TrimSpace(strings.TrimPrefix(line , "Active Profile: ")) , nil
		}
	}

	return "" , nil
}