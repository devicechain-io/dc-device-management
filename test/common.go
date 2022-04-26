/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package test

import (
	"context"
	"time"

	"github.com/devicechain-io/dc-device-management/model"
	"github.com/devicechain-io/dc-microservice/config"
	"github.com/devicechain-io/dc-microservice/core"
	"github.com/stretchr/testify/mock"
)

// Microservice information used for testing.
var DeviceManagementMicroservice = &core.Microservice{
	StartTime:                    time.Now(),
	InstanceId:                   "devicechain",
	TenantId:                     "tenant1",
	TenantName:                   "Tenant 1",
	MicroserviceId:               "device-management",
	MicroserviceName:             "Device Management",
	FunctionalArea:               "device-management",
	InstanceConfiguration:        config.InstanceConfiguration{},
	MicroserviceConfigurationRaw: make([]byte, 0),
}

/**
 * Mock for device management API.
 */

type MockApi struct {
	mock.Mock
}

func (api *MockApi) DeviceTypeById(ctx context.Context, id uint) (*model.DeviceType, error) {
	args := api.Mock.Called()
	return args.Get(0).(*model.DeviceType), args.Error(1)
}

func (api *MockApi) DeviceTypeByToken(ctx context.Context, token string) (*model.DeviceType, error) {
	args := api.Mock.Called()
	return args.Get(0).(*model.DeviceType), args.Error(1)
}

func (api *MockApi) DeviceTypes(ctx context.Context, criteria model.DeviceTypeSearchCriteria) (*model.DeviceTypeSearchResults, error) {
	args := api.Mock.Called()
	return args.Get(0).(*model.DeviceTypeSearchResults), args.Error(1)
}

func (api *MockApi) DeviceById(ctx context.Context, id uint) (*model.Device, error) {
	args := api.Mock.Called()
	return args.Get(0).(*model.Device), args.Error(1)
}

func (api *MockApi) DeviceByToken(ctx context.Context, token string) (*model.Device, error) {
	args := api.Mock.Called()
	return args.Get(0).(*model.Device), args.Error(1)
}

func (api *MockApi) Devices(ctx context.Context, criteria model.DeviceSearchCriteria) (*model.DeviceSearchResults, error) {
	args := api.Mock.Called()
	return args.Get(0).(*model.DeviceSearchResults), args.Error(1)
}

func (api *MockApi) CreateDeviceAssignment(ctx context.Context,
	request *model.DeviceAssignmentCreateRequest) (*model.DeviceAssignment, error) {
	args := api.Mock.Called()
	return args.Get(0).(*model.DeviceAssignment), args.Error(1)
}

func (api *MockApi) ActiveDeviceAssignmentsForDevice(ctx context.Context, id uint) ([]model.DeviceAssignment, error) {
	args := api.Mock.Called()
	return args.Get(0).([]model.DeviceAssignment), args.Error(1)
}
