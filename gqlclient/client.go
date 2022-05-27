/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

// The gqlclient package provides a GraphQL client for commonly used operations
// in the device management API.
package gqlclient

//go:generate go run github.com/Khan/genqlient@v0.4.0

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/devicechain-io/dc-device-management/model"
)

// TEMP: genqlient doesn't handle pointers yet.
func blank(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

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

// Get a device type by id.
func GetDeviceTypeById(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*getDeviceTypeByIdResponse, error) {
	return getDeviceTypeById(ctx, client, id)
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
