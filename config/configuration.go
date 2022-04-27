/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package config

import (
	"github.com/devicechain-io/dc-microservice/config"
)

const (
	KAFKA_TOPIC_FAILED_EVENTS   = "failed-events"
	KAFKA_TOPIC_RESOLVED_EVENTS = "resolved-events"
)

type DeviceManagementConfiguration struct {
	RdbConfiguration config.MicroserviceRdbConfiguration
}

// Creates the default device management configuration
func NewDeviceManagementConfiguration() *DeviceManagementConfiguration {
	return &DeviceManagementConfiguration{
		RdbConfiguration: config.MicroserviceRdbConfiguration{
			SqlDebug: true,
		},
	}
}
