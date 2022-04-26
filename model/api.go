/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"context"

	"github.com/devicechain-io/dc-microservice/rdb"
)

type Api struct {
	RDB *rdb.RdbManager
}

// Create a new API instance.
func NewApi(rdb *rdb.RdbManager) *Api {
	api := &Api{}
	api.RDB = rdb
	return api
}

// Interface for device management API (used for mocking)
type DeviceManagementApi interface {
	// Device types.
	DeviceTypeById(ctx context.Context, id uint) (*DeviceType, error)
	DeviceTypeByToken(ctx context.Context, token string) (*DeviceType, error)
	DeviceTypes(ctx context.Context, criteria DeviceTypeSearchCriteria) (*DeviceTypeSearchResults, error)

	// Devices.
	DeviceById(ctx context.Context, id uint) (*Device, error)
	DeviceByToken(ctx context.Context, token string) (*Device, error)
	Devices(ctx context.Context, criteria DeviceSearchCriteria) (*DeviceSearchResults, error)

	// Device assignments.
	CreateDeviceAssignment(ctx context.Context, request *DeviceAssignmentCreateRequest) (*DeviceAssignment, error)
	ActiveDeviceAssignmentsForDevice(ctx context.Context, id uint) ([]DeviceAssignment, error)
}
