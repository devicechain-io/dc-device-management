/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

// The gqlclient package provides a GraphQL client for commonly used operations
// in the device management API.
package gqlclient

import "github.com/devicechain-io/dc-device-management/model"

//go:generate go run github.com/Khan/genqlient@v0.5.0

// Converts relationship targets into generated datatype.
func targets(req model.EntityRelationshipCreateRequest) EntityRelationshipTargetsCreateRequest {
	return EntityRelationshipTargetsCreateRequest{
		TargetDevice:        req.TargetDevice,
		TargetDeviceGroup:   req.TargetDeviceGroup,
		TargetAsset:         req.TargetAsset,
		TargetAssetGroup:    req.TargetAssetGroup,
		TargetArea:          req.TargetArea,
		TargetAreaGroup:     req.TargetAreaGroup,
		TargetCustomer:      req.TargetCustomer,
		TargetCustomerGroup: req.TargetCustomerGroup,
	}
}
