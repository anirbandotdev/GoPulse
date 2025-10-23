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

func GetBluetoothCard() (string, error) {
	cmd := exec.Command("pactl", "list", "short", "cards")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "bluez_card.") {
			fields := strings.Fields(line)
			if len(fields) > 1 {
				return fields[1], nil
			}
		}
	}
	return "", nil
}


var profileAliases = map[string]string{
	"music": "a2dp-sink",
	"call":  "headset-head-unit",
	"a2dp":  "a2dp-sink",
	"hsp":   "headset-head-unit",
	"hfp":   "headset-head-unit",
}

func ResolveProfileAlias(input string) string {
	if real, ok := profileAliases[strings.ToLower(input)]; ok {
		return real
	}
	return input 
}
