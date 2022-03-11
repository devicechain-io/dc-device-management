/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package main

import (
	"fmt"

	"github.com/devicechain-io/dc-microservice/config"
)

type NestedConfiguration struct {
	Test string
}

type DeviceManagementConfiguration struct {
	Nested NestedConfiguration
}

// Get instance configuration CRs that should be created in tooling
func GetInstanceConfigurationResources() ([]config.ConfigurationResource, error) {
	resources := make([]config.ConfigurationResource, 0)

	name := "device-management-default"
	msconfig := NewDeviceManagementConfiguration()
	content, err := config.GenerateMicroserviceConfig(name, msconfig)
	if err != nil {
		return nil, err
	}
	dcmdefault := config.ConfigurationResource{
		Name:    fmt.Sprintf("%s_%s", "core.devicechain.io", name),
		Content: content,
	}

	resources = append(resources, dcmdefault)
	return resources, nil
}

// Creates the default device management configuration
func NewDeviceManagementConfiguration() *DeviceManagementConfiguration {
	return &DeviceManagementConfiguration{
		Nested: NestedConfiguration{
			Test: "test",
		},
	}
}
