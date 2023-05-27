// Copyright Tharsis Labs Ltd.(Gridiron)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/fury-labs/blackfury/blob/main/LICENSE)

package v8

const (
	// UpgradeName is the shared upgrade plan name for mainnet and testnet
	UpgradeName = "v8.0.0"
	// MainnetUpgradeHeight defines the Gridiron mainnet block height on which the upgrade will take place
	MainnetUpgradeHeight = 3_489_000
	// TestnetUpgradeHeight defines the Gridiron testnet block height on which the upgrade will take place
	TestnetUpgradeHeight = 4_600_000
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/fury-labs/blackfury/releases/download/v8.0.0/blackfury_8.0.0_Darwin_arm64.tar.gz","darwin/x86_64":"https://github.com/fury-labs/blackfury/releases/download/v8.0.0/blackfury_8.0.0_Darwin_x86_64.tar.gz","linux/arm64":"https://github.com/fury-labs/blackfury/releases/download/v8.0.0/blackfury_8.0.0_Linux_arm64.tar.gz","linux/amd64":"https://github.com/fury-labs/blackfury/releases/download/v8.0.0/blackfury_8.0.0_Linux_x86_64.tar.gz","windows/x86_64":"https://github.com/fury-labs/blackfury/releases/download/v8.0.0/blackfury_8.0.0_Windows_x86_64.zip"}}'`
)
