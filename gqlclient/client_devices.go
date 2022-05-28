/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

// The gqlclient package provides a GraphQL client for commonly used operations
// in the device management API.
package gqlclient

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/devicechain-io/dc-device-management/model"
)

// Assure that a device type exists.
func AssureDeviceType(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceTypeCreateRequest,
) (*getDeviceTypeByTokenResponse, *createDeviceTypeResponse, error) {
	gresp, err := GetDeviceTypeByToken(ctx, client, request.Token)
	if err == nil {
		return gresp, nil, nil
	}
	cresp, err := CreateDeviceType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new device type.
func CreateDeviceType(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceTypeCreateRequest,
) (*createDeviceTypeResponse, error) {
	return createDeviceType(ctx, client, request.Token, blank(request.Name), blank(request.Description),
		blank(request.ImageUrl), blank(request.Icon), blank(request.BackgroundColor), blank(request.ForegroundColor),
		blank(request.BorderColor), blank(request.Metadata))
}

// Get a device type by token.
func GetDeviceTypeByToken(
	ctx context.Context,
	client graphql.Client,
	token string,
) (*getDeviceTypeByTokenResponse, error) {
	return getDeviceTypeByToken(ctx, client, token)
}

// List device types based on criteria.
func ListDeviceTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listDeviceTypesResponse, error) {
	return listDeviceTypes(ctx, client, pageNumber, pageSize)
}

// Assure that a device exists.
func AssureDevice(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceCreateRequest,
) (*getDeviceByTokenResponse, *createDeviceResponse, error) {
	gresp, err := GetDeviceByToken(ctx, client, request.Token)
	if err == nil {
		return gresp, nil, nil
	}
	cresp, err := CreateDevice(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new device.
func CreateDevice(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceCreateRequest,
) (*createDeviceResponse, error) {
	return createDevice(ctx, client, request.Token, request.DeviceTypeToken,
		blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get a device by token.
func GetDeviceByToken(
	ctx context.Context,
	client graphql.Client,
	token string,
) (*getDeviceByTokenResponse, error) {
	return getDeviceByToken(ctx, client, token)
}

// List devices based on criteria.
func ListDevices(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listDevicesResponse, error) {
	return listDevices(ctx, client, pageNumber, pageSize)
}

// Assure that a device relationship type exists.
func AssureDeviceRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceRelationshipTypeCreateRequest,
) (*getDeviceRelationshipTypeByTokenResponse, *createDeviceRelationshipTypeResponse, error) {
	gresp, err := GetDeviceRelationshipTypeByToken(ctx, client, request.Token)
	if err == nil {
		return gresp, nil, nil
	}
	cresp, err := CreateDeviceRelationshipType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new device relationship type.
func CreateDeviceRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceRelationshipTypeCreateRequest,
) (*createDeviceRelationshipTypeResponse, error) {
	return createDeviceRelationshipType(ctx, client, request.Token, blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get a device relationship type by token.
func GetDeviceRelationshipTypeByToken(
	ctx context.Context,
	client graphql.Client,
	token string,
) (*getDeviceRelationshipTypeByTokenResponse, error) {
	return getDeviceRelationshipTypeByToken(ctx, client, token)
}

// List device relationship types based on criteria.
func ListDeviceRelationshipTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listDeviceRelationshipTypesResponse, error) {
	return listDeviceRelationshipTypes(ctx, client, pageNumber, pageSize)
}

// Assure that a device relationship exists.
func AssureDeviceRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceRelationshipCreateRequest,
) (*getDeviceRelationshipByTokenResponse, *createDeviceRelationshipResponse, error) {
	gresp, err := GetDeviceRelationshipByToken(ctx, client, request.Token)
	if err == nil {
		return gresp, nil, nil
	}
	cresp, err := CreateDeviceRelationship(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new device relationship.
func CreateDeviceRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceRelationshipCreateRequest,
) (*createDeviceRelationshipResponse, error) {
	return createDeviceRelationship(ctx, client, request.Token, request.SourceDevice, request.TargetDevice, request.RelationshipType)
}

// Get a device relationship by token.
func GetDeviceRelationshipByToken(
	ctx context.Context,
	client graphql.Client,
	token string,
) (*getDeviceRelationshipByTokenResponse, error) {
	return getDeviceRelationshipByToken(ctx, client, token)
}

// List device relationships based on criteria.
func ListDeviceRelationships(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listDeviceRelationshipsResponse, error) {
	return listDeviceRelationships(ctx, client, pageNumber, pageSize)
}

// Assure that a device group exists.
func AssureDeviceGroup(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceGroupCreateRequest,
) (*getDeviceGroupByTokenResponse, *createDeviceGroupResponse, error) {
	gresp, err := GetDeviceGroupByToken(ctx, client, request.Token)
	if err == nil {
		return gresp, nil, nil
	}
	cresp, err := CreateDeviceGroup(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new device group.
func CreateDeviceGroup(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceGroupCreateRequest,
) (*createDeviceGroupResponse, error) {
	return createDeviceGroup(ctx, client, request.Token, blank(request.Name), blank(request.Description),
		blank(request.ImageUrl), blank(request.Icon), blank(request.BackgroundColor), blank(request.ForegroundColor),
		blank(request.BorderColor), blank(request.Metadata))
}

// Get a device group by token.
func GetDeviceGroupByToken(
	ctx context.Context,
	client graphql.Client,
	token string,
) (*getDeviceGroupByTokenResponse, error) {
	return getDeviceGroupByToken(ctx, client, token)
}

// List device groups based on criteria.
func ListDeviceGroups(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listDeviceGroupsResponse, error) {
	return listDeviceGroups(ctx, client, pageNumber, pageSize)
}

// Assure that a device group relationship type exists.
func AssureDeviceGroupRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceGroupRelationshipTypeCreateRequest,
) (*getDeviceGroupRelationshipTypeByTokenResponse, *createDeviceGroupRelationshipTypeResponse, error) {
	gresp, err := GetDeviceGroupRelationshipTypeByToken(ctx, client, request.Token)
	if err == nil {
		return gresp, nil, nil
	}
	cresp, err := CreateDeviceGroupRelationshipType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new device group relationship type.
func CreateDeviceGroupRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceGroupRelationshipTypeCreateRequest,
) (*createDeviceGroupRelationshipTypeResponse, error) {
	return createDeviceGroupRelationshipType(ctx, client, request.Token, blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get a device group relationship type by token.
func GetDeviceGroupRelationshipTypeByToken(
	ctx context.Context,
	client graphql.Client,
	token string,
) (*getDeviceGroupRelationshipTypeByTokenResponse, error) {
	return getDeviceGroupRelationshipTypeByToken(ctx, client, token)
}

// List device group relationship types based on criteria.
func ListDeviceGroupRelationshipTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listDeviceGroupRelationshipTypesResponse, error) {
	return listDeviceGroupRelationshipTypes(ctx, client, pageNumber, pageSize)
}

// Assure that a device group relationship exists.
func AssureDeviceGroupRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceGroupRelationshipCreateRequest,
) (*getDeviceGroupRelationshipByTokenResponse, *createDeviceGroupRelationshipResponse, error) {
	gresp, err := GetDeviceGroupRelationshipByToken(ctx, client, request.Token)
	if err == nil {
		return gresp, nil, nil
	}
	cresp, err := CreateDeviceGroupRelationship(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new device group relationship.
func CreateDeviceGroupRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.DeviceGroupRelationshipCreateRequest,
) (*createDeviceGroupRelationshipResponse, error) {
	return createDeviceGroupRelationship(ctx, client, request.Token, request.DeviceGroup, request.Device, request.RelationshipType)
}

// Get a device group relationship by token.
func GetDeviceGroupRelationshipByToken(
	ctx context.Context,
	client graphql.Client,
	token string,
) (*getDeviceGroupRelationshipByTokenResponse, error) {
	return getDeviceGroupRelationshipByToken(ctx, client, token)
}

// List device group relationships based on criteria.
func ListDeviceGroupRelationships(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listDeviceGroupRelationshipsResponse, error) {
	return listDeviceGroupRelationships(ctx, client, pageNumber, pageSize)
}
