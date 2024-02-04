package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"firewall-config-checker/internal/firewall"
)

var rootCmd = &cobra.Command{
	Use:   "firewallcheck",
	Short: "Check and toggle firewall configuration",
	Run: func(cmd *cobra.Command, args []string) {
		checkAndScan()
	},
}

var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle firewall state",
	Run: func(cmd *cobra.Command, args []string) {
		toggleFirewall()
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func checkAndScan() {
	// Get user input for target (localhost or specific IP address)
	fmt.Print("Enter the target (localhost or specific IP address): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	target := scanner.Text()

	// Perform port scanning for the provided target
	portScanResult, err := performPortScan(target)
	if err != nil {
		log.Fatal("Error performing port scan:", err)
	}

	fmt.Println("Port Scan Result for", target, ":\n", portScanResult)

	// Ask the user if they want to expand scanning to the LAN or other addresses
	fmt.Print("Do you want to expand scanning to the LAN or specific addresses? (y/n): ")
	scanner.Scan()
	response := strings.ToLower(scanner.Text())
	if response == "y" {
		// Implement logic to expand scanning to the LAN or specific addresses
		// You can use additional functions or external libraries for more advanced scanning
		fmt.Println("Expanding scanning to LAN or specific addresses...")
		// Example: firewall.ScanLAN(target)
	} else {
		fmt.Println("Exiting.")
	}
}

func performPortScan(target string) (string, error) {
	cmd := exec.Command("nmap", "-p-", target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func toggleFirewall() {
	firewallStatus, err := firewall.GetFirewallStatus()
	if err != nil {
		fmt.Println("Error checking firewall:", err)
		return
	}

	fmt.Println("Current Firewall Status:", firewallStatus)

	if strings.Contains(firewallStatus, "ON") {
		fmt.Println("Turning off the firewall...")
		err = firewall.SetFirewallState("off")
	} else {
		fmt.Println("Turning on the firewall...")
		err = firewall.SetFirewallState("on")
	}

	if err != nil {
		fmt.Println("Error toggling firewall state:", err)
		return
	}

	newFirewallStatus, err := firewall.GetFirewallStatus()
	if err != nil {
		fmt.Println("Error checking firewall after toggle:", err)
		return
	}

	fmt.Println("Firewall Status after Toggle:", newFirewallStatus)
}
