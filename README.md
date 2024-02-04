# Security Tool - Firewall Configuration Checker and Digital Forensics

## Firewall Configuration Checker

This module automates firewall checks, identifies vulnerabilities through CVE scanning, and uses Colly to crawl the LAN network for targeted security checks.

**Features:**
- **Automated Firewall Checks:** Quickly assess the configuration of your firewall settings.
- **CVE Scanning:** Identify potential vulnerabilities through CVE scanning.
- **LAN Network Crawler:** Utilize Colly to discover devices on the LAN network.

## Digital Forensics

This module allows users to choose between analyzing victim and perpetrator systems, selecting from various victim system categories and performing specific forensics tasks.

**Features:**
- **Victim and Perpetrator System Analysis:** Choose between analyzing victim and perpetrator systems before selecting a specific forensic task.
- **Victim System Categories:**
  - Servers hacked or cracked through vulnerabilities and bad server configuration.
  - Unauthorized access to servers through stolen credentials.
  - Computers attacked through malware by downloading malicious files or scripts.
- **Forensic Task Options:** Perform various forensics tasks, including network forensics, memory forensics, disk forensics, and incident response automation.

## Automated Victim System Analysis

This forensics tool is designed to assist users in investigating compromised victim systems. It provides an automated analysis option to identify vulnerabilities, anomalies, and potential security incidents.

### Getting Started

**Prerequisites:**
- Go 1.21.6 or higher
- Dependencies specified in `go.mod`

### Installation

```bash
go run main.go
