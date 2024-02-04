# Security Tool - Firewall Configuration Checker and Digital Forensics

## Firewall Configuration Checker

### Overview

<p>This module provides automated checks for your firewall settings and identifies potential vulnerabilities through CVE scanning. Additionally, it utilizes Colly to crawl the LAN network, discovering devices for targeted security checks.</p>

### Features

- <h3>Automated Firewall Checks:</h3> <p>Quickly assess the configuration of your firewall settings.</p>
- <h3>CVE Scanning:</h3> <p>Identify potential vulnerabilities by cross-referencing service versions with known Common Vulnerabilities and Exposures (CVEs).</p>
- <h3>LAN Network Crawler:</h3> <p>Use Colly to crawl the LAN network and discover devices for targeted security checks.</p>

## Digital Forensics

### Overview

<p>This module allows users to choose between analyzing victim and perpetrator systems before selecting a specific forensic task. Users can select from various victim system categories and perform specific forensics tasks.</p>

### Features

- <h3>Victim and Perpetrator System Analysis:</h3> <p>Choose between analyzing victim and perpetrator systems before selecting a specific forensic task.</p>
- <h3>Victim System Categories:</h3>
  - <p>Servers that have been hacked or cracked through system vulnerabilities and bad server configuration.</p>
  - <p>Unauthorized access cases to servers through stolen credentials.</p>
  - <p>Computers that have been attacked through malware by downloading malicious files or scripts.</p>
- <h3>Forensic Task Options:</h3> <p>Perform various forensics tasks, including network forensics, memory forensics, disk forensics, and incident response automation.</p>

## Getting Started

### Prerequisites

- Go 1.21.6 or higher

### Installation

```bash
go run main.go
