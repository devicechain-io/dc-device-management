/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package config

type NestedConfiguration struct {
	Test string
}

type DeviceManagementConfiguration struct {
	Nested NestedConfiguration
}

// Resource provider
type ResourceProvider struct {
}

// Creates the default device management configuration
func NewDeviceManagementConfiguration() *DeviceManagementConfiguration {
	return &DeviceManagementConfiguration{
		Nested: NestedConfiguration{
			Test: "test",
		},
	}
}
