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
	Metadata        *string
}

// Represents a device type.
type DeviceType struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.BrandedEntity
	rdb.MetadataEntity

	Devices []Device
}

// Search criteria for locating device types.
type DeviceTypeSearchCriteria struct {
	rdb.Pagination
}

// Data required to create a device.
type DeviceCreateRequest struct {
	Token           string
	Name            *string
	Description     *string
	DeviceTypeToken string
	Metadata        *string
}

// Represents a device.
type Device struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.MetadataEntity

	DeviceTypeId int
	DeviceType   *DeviceType
}

// Search criteria for locating devices.
type DeviceSearchCriteria struct {
	rdb.Pagination
	DeviceTypeToken *string
}

// Data required to create a device relationship type.
type DeviceRelationshipTypeCreateRequest struct {
	Token       string
	Name        *string
	Description *string
	Metadata    *string
}

// Metadata indicating a relationship between devices.
type DeviceRelationshipType struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.MetadataEntity
}

// Search criteria for locating device relationship types.
type DeviceRelationshipTypeSearchCriteria struct {
	rdb.Pagination
}

// Data required to create a device relationship.
type DeviceRelationshipCreateRequest struct {
	SourceDevice     string
	TargetDevice     string
	RelationshipType string
	Metadata         *string
}

// Captures a relationship between devices.
type DeviceRelationship struct {
	gorm.Model
	rdb.MetadataEntity
	SourceDeviceId     int
	SourceDevice       Device
	TargetDeviceId     int
	TargetDevice       Device
	RelationshipTypeId int
	RelationshipType   DeviceRelationshipType
}

// Search criteria for locating device relationships.
type DeviceRelationshipSearchCriteria struct {
	rdb.Pagination
}

// Data required to create a device group.
type DeviceGroupCreateRequest struct {
	Token           string
	Name            *string
	Description     *string
	ImageUrl        *string
	Icon            *string
	BackgroundColor *string
	ForegroundColor *string
	BorderColor     *string
	Metadata        *string
}

// Represents a group of devices.
type DeviceGroup struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.BrandedEntity
	rdb.MetadataEntity
}

// Search criteria for locating device groups.
type DeviceGroupSearchCriteria struct {
	rdb.Pagination
}

// Data required to create a device group relationship type.
type DeviceGroupRelationshipTypeCreateRequest struct {
	Token       string
	Name        *string
	Description *string
	Metadata    *string
}

// Metadata indicating a relationship between device and group.
type DeviceGroupRelationshipType struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.MetadataEntity
}

// Search criteria for locating device groups relationship types.
type DeviceGroupRelationshipTypeSearchCriteria struct {
	rdb.Pagination
}

// Data required to create a device group relationship.
type DeviceGroupRelationshipCreateRequest struct {
	DeviceGroup      string
	Device           string
	RelationshipType string
	Metadata         *string
}

// Represents a device-to-group relationship.
type DeviceGroupRelationship struct {
	gorm.Model
	rdb.MetadataEntity
	DeviceGroupId      int
	DeviceGroup        DeviceGroup
	DeviceId           int
	Device             Device
	RelationshipTypeId int
	RelationshipType   DeviceGroupRelationshipType
}

// Search criteria for locating device groups relationships.
type DeviceGroupRelationshipSearchCriteria struct {
	rdb.Pagination
}
