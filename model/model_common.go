/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"github.com/devicechain-io/dc-microservice/rdb"
	"gorm.io/gorm"
)

// Base data required to create an entity relationship.
type EntityRelationshipCreateRequest struct {
	TargetDevice        *string
	TargetDeviceGroup   *string
	TargetAsset         *string
	TargetAssetGroup    *string
	TargetArea          *string
	TargetAreaGroup     *string
	TargetCustomer      *string
	TargetCustomerGroup *string
}

// Based data for capturing a relationship between entites.
type EntityRelationship struct {
	gorm.Model
	rdb.TokenReference
	rdb.MetadataEntity
	TargetDeviceId        *uint
	TargetDevice          *Device
	TargetDeviceGroupId   *uint
	TargetDeviceGroup     *DeviceGroup
	TargetAssetId         *uint
	TargetAsset           *Asset
	TargetAssetGroupId    *uint
	TargetAssetGroup      *AssetGroup
	TargetAreaId          *uint
	TargetArea            *Area
	TargetAreaGroupId     *uint
	TargetAreaGroup       *AreaGroup
	TargetCustomerId      *uint
	TargetCustomer        *Customer
	TargetCustomerGroupId *uint
	TargetCustomerGroup   *CustomerGroup
}
