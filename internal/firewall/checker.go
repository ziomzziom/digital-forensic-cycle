package firewall

import (
	"os/exec"
	"strings"
)

// GetFirewallStatus retrieves the firewall status using netsh command
func GetFirewallStatus() (string, error) {
	cmd := exec.Command("netsh", "advfirewall", "show", "allprofiles", "state")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

// SetFirewallState sets the firewall state using netsh command
func SetFirewallState(state string) error {
	cmd := exec.Command("netsh", "advfirewall", "set", "allprofiles", "state", state)
	err := cmd.Run()
	return err
}
