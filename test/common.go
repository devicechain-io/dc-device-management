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

func (api *MockApi) DeviceTypesById(ctx context.Context, ids []uint) ([]*model.DeviceType, error) {
	args := api.Mock.Called()
	return args.Get(0).([]*model.DeviceType), args.Error(1)
}

func (api *MockApi) DeviceTypesByToken(ctx context.Context, tokens []string) ([]*model.DeviceType, error) {
	args := api.Mock.Called()
	return args.Get(0).([]*model.DeviceType), args.Error(1)
}

func (api *MockApi) DeviceTypes(ctx context.Context, criteria model.DeviceTypeSearchCriteria) (*model.DeviceTypeSearchResults, error) {
	args := api.Mock.Called()
	return args.Get(0).(*model.DeviceTypeSearchResults), args.Error(1)
}

func (api *MockApi) DevicesById(ctx context.Context, ids []uint) ([]*model.Device, error) {
	args := api.Mock.Called()
	return args.Get(0).([]*model.Device), args.Error(1)
}

func (api *MockApi) DevicesByToken(ctx context.Context, tokens []string) ([]*model.Device, error) {
	args := api.Mock.Called()
	return args.Get(0).([]*model.Device), args.Error(1)
}

func (api *MockApi) Devices(ctx context.Context, criteria model.DeviceSearchCriteria) (*model.DeviceSearchResults, error) {
	args := api.Mock.Called()
	return args.Get(0).(*model.DeviceSearchResults), args.Error(1)
}

func (api *MockApi) DeviceRelationshipsById(ctx context.Context, ids []uint) ([]*model.DeviceRelationship, error) {
	args := api.Mock.Called()
	return args.Get(0).([]*model.DeviceRelationship), args.Error(1)
}

func (api *MockApi) DeviceRelationshipsByToken(ctx context.Context, tokens []string) ([]*model.DeviceRelationship, error) {
	args := api.Mock.Called()
	return args.Get(0).([]*model.DeviceRelationship), args.Error(1)
}

func (api *MockApi) DeviceRelationships(ctx context.Context,
	criteria model.DeviceRelationshipSearchCriteria) (*model.DeviceRelationshipSearchResults, error) {
	args := api.Mock.Called()
	return args.Get(0).(*model.DeviceRelationshipSearchResults), args.Error(1)
}

func (api *MockApi) CreateDeviceRelationship(ctx context.Context, request *model.DeviceRelationshipCreateRequest) (*model.DeviceRelationship, error) {
	args := api.Mock.Called()
	return args.Get(0).(*model.DeviceRelationship), args.Error(1)
}
