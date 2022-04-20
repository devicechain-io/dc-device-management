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

// Data required to create a device assignment status.
type DeviceAssignmentStatusCreateRequest struct {
	Token       string
	Name        *string
	Description *string
	Metadata    *string
}

// Allows status to be associated with a device assignment.
type DeviceAssignmentStatus struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.MetadataEntity
}

// Search criteria for locating assignment statuses.
type DeviceAssignmentStatusSearchCriteria struct {
	rdb.Pagination
}

// Results for assignment status search.
type DeviceAssignmentStatusSearchResults struct {
	Results    []DeviceAssignmentStatus
	Pagination rdb.SearchResultsPagination
}

// Data required to create a device assignment.
type DeviceAssignmentCreateRequest struct {
	Token                  string
	Device                 string
	DeviceGroup            *string
	Asset                  *string
	AssetGroup             *string
	Customer               *string
	CustomerGroup          *string
	Area                   *string
	AreaGroup              *string
	Metadata               *string
	DeviceAssignmentStatus *string
	Active                 *bool
}

// Provides context for device.
type DeviceAssignment struct {
	gorm.Model
	rdb.TokenReference
	rdb.MetadataEntity
	DeviceId                 uint
	Device                   Device
	DeviceGroupId            *uint
	DeviceGroup              *DeviceGroup
	AssetId                  *uint
	Asset                    *Asset
	AssetGroupId             *uint
	AssetGroup               *AssetGroup
	CustomerId               *uint
	Customer                 *Customer
	CustomerGroupId          *uint
	CustomerGroup            *CustomerGroup
	AreaId                   *uint
	Area                     *Area
	AreaGroupId              *uint
	AreaGroup                *AreaGroup
	DeviceAssignmentStatusId *uint
	DeviceAssignmentStatus   *DeviceAssignmentStatus
	Active                   bool
}

// Search criteria for locating device assignments.
type DeviceAssignmentSearchCriteria struct {
	rdb.Pagination
	Device      *string
	DeviceGroup *string
}

// Results for device assignment search.
type DeviceAssignmentSearchResults struct {
	Results    []DeviceAssignment
	Pagination rdb.SearchResultsPagination
}
