/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

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
) (*getDeviceTypesByTokenResponse, *createDeviceTypeResponse, error) {
	gresp, err := GetDeviceTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.DeviceTypesByToken) > 0 {
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

// Get device types by token.
func GetDeviceTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getDeviceTypesByTokenResponse, error) {
	return getDeviceTypesByToken(ctx, client, tokens)
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
) (*getDevicesByTokenResponse, *createDeviceResponse, error) {
	gresp, err := GetDevicesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.DevicesByToken) > 0 {
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

// Get devices by token.
func GetDevicesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getDevicesByTokenResponse, error) {
	return getDevicesByToken(ctx, client, tokens)
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
) (*getDeviceRelationshipTypesByTokenResponse, *createDeviceRelationshipTypeResponse, error) {
	gresp, err := GetDeviceRelationshipTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.DeviceRelationshipTypesByToken) > 0 {
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

// Get device relationship types by token.
func GetDeviceRelationshipTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getDeviceRelationshipTypesByTokenResponse, error) {
	return getDeviceRelationshipTypesByToken(ctx, client, tokens)
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
) (*getDeviceRelationshipsByTokenResponse, *createDeviceRelationshipResponse, error) {
	gresp, err := GetDeviceRelationshipsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.DeviceRelationshipsByToken) > 0 {
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

// Get device relationships by token.
func GetDeviceRelationshipsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getDeviceRelationshipsByTokenResponse, error) {
	return getDeviceRelationshipsByToken(ctx, client, tokens)
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
) (*getDeviceGroupsByTokenResponse, *createDeviceGroupResponse, error) {
	gresp, err := GetDeviceGroupsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.DeviceGroupsByToken) > 0 {
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

// Get device groups by token.
func GetDeviceGroupsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getDeviceGroupsByTokenResponse, error) {
	return getDeviceGroupsByToken(ctx, client, tokens)
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
) (*getDeviceGroupRelationshipTypesByTokenResponse, *createDeviceGroupRelationshipTypeResponse, error) {
	gresp, err := GetDeviceGroupRelationshipTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.DeviceGroupRelationshipTypesByToken) > 0 {
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

// Get a device group relationship types by token.
func GetDeviceGroupRelationshipTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getDeviceGroupRelationshipTypesByTokenResponse, error) {
	return getDeviceGroupRelationshipTypesByToken(ctx, client, tokens)
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
) (*getDeviceGroupRelationshipsByTokenResponse, *createDeviceGroupRelationshipResponse, error) {
	gresp, err := GetDeviceGroupRelationshipsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.DeviceGroupRelationshipsByToken) > 0 {
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

// Get a device group relationships by token.
func GetDeviceGroupRelationshipsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getDeviceGroupRelationshipsByTokenResponse, error) {
	return getDeviceGroupRelationshipsByToken(ctx, client, tokens)
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
