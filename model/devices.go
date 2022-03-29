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

// Represents a device type.
type DeviceType struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.BrandedEntity

	Devices []Device
}

// Search criteria for locating device types.
type DeviceTypeSearchCriteria struct {
	rdb.Pagination
}

// Data required to create a device type.
type DeviceTypeCreateRequest struct {
	Token           string
	Name            *string
	Description     *string
	ImageUrl        *string
	Icon            *string
	BackgroundColor *string
	ForegroundColor *string
	BorderColor     *string
}

// Represents a device.
type Device struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity

	DeviceTypeId int
	DeviceType   DeviceType
}

// Data required to create a device.
type DeviceCreateRequest struct {
	Token           string
	Name            *string
	Description     *string
	DeviceTypeToken string
}

// Search criteria for locating devices.
type DeviceSearchCriteria struct {
	rdb.Pagination
	DeviceTypeToken *string
}

// Metadata indicating a relationship between devices.
type DeviceRelationshipType struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
}

// Captures a relationship between devices.
type DeviceRelationship struct {
	gorm.Model

	SourceDeviceId     int
	SourceDevice       Device
	TargetDeviceId     int
	TargetDevice       Device
	RelationshipTypeId int
	RelationshipType   DeviceRelationshipType
}

// Represents a group of devices.
type DeviceGroup struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.BrandedEntity
}

// Role of device within a group.
type DeviceGroupRole struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
}

// Represents a device-to-group relationship.
type DeviceGroupRelationship struct {
	gorm.Model

	DeviceGroupId     int
	DeviceGroup       DeviceGroup
	DeviceId          int
	Device            Device
	DeviceGroupRoleId int
	DeviceGroupRole   DeviceGroupRole
}

// Role of subgroup within a group.
type DeviceSubroupRole struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
}

// Represents a subgroup-to-group relationship.
type DeviceSubgroupRelationship struct {
	gorm.Model

	DeviceGroupId       int
	DeviceGroup         DeviceGroup
	SubgroupId          int
	Subgroup            DeviceGroup
	DeviceSubroupRoleId int
	DeviceSubroupRole   DeviceSubroupRole
}
