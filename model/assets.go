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

// Data required to create an asset type.
type AssetTypeCreateRequest struct {
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

// Represents an asset type.
type AssetType struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.BrandedEntity
	rdb.MetadataEntity

	Assets []Asset
}

// Search criteria for locating asset types.
type AssetTypeSearchCriteria struct {
	rdb.Pagination
}

// Data required to create an asset.
type AssetCreateRequest struct {
	Token          string
	Name           *string
	Description    *string
	AssetTypeToken string
	Metadata       *string
}

// Represents an asset.
type Asset struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.MetadataEntity

	AssetTypeId int
	AssetType   *AssetType
}

// Search criteria for locating assets.
type AssetSearchCriteria struct {
	rdb.Pagination
	AssetTypeToken *string
}

// Data required to create an asset relationship type.
type AssetRelationshipTypeCreateRequest struct {
	Token       string
	Name        *string
	Description *string
	Metadata    *string
}

// Metadata indicating a relationship between assets.
type AssetRelationshipType struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.MetadataEntity
}

// Search criteria for locating asset relationship types.
type AssetRelationshipTypeSearchCriteria struct {
	rdb.Pagination
}

// Data required to create an asset relationship.
type AssetRelationshipCreateRequest struct {
	SourceAsset      string
	TargetAsset      string
	RelationshipType string
	Metadata         *string
}

// Captures a relationship between assets.
type AssetRelationship struct {
	gorm.Model
	rdb.MetadataEntity
	SourceAssetId      int
	SourceAsset        Asset
	TargetAssetId      int
	TargetAsset        Asset
	RelationshipTypeId int
	RelationshipType   AssetRelationshipType
}

// Search criteria for locating asset relationships.
type AssetRelationshipSearchCriteria struct {
	rdb.Pagination
}

// Data required to create an asset group.
type AssetGroupCreateRequest struct {
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

// Represents a group of assets.
type AssetGroup struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.BrandedEntity
	rdb.MetadataEntity
}

// Search criteria for locating asset groups.
type AssetGroupSearchCriteria struct {
	rdb.Pagination
}

// Data required to create an asset group relationship type.
type AssetGroupRelationshipTypeCreateRequest struct {
	Token       string
	Name        *string
	Description *string
	Metadata    *string
}

// Metadata indicating a relationship between asset and group.
type AssetGroupRelationshipType struct {
	gorm.Model
	rdb.TokenReference
	rdb.NamedEntity
	rdb.MetadataEntity
}

// Search criteria for locating asset groups relationship types.
type AssetGroupRelationshipTypeSearchCriteria struct {
	rdb.Pagination
}

// Data required to create an asset group relationship.
type AssetGroupRelationshipCreateRequest struct {
	DeviceGroup      string
	Device           string
	RelationshipType string
	Metadata         *string
}

// Represents a asset-to-group relationship.
type AssetGroupRelationship struct {
	gorm.Model
	rdb.MetadataEntity
	AssetGroupId       int
	AssetGroup         AssetGroup
	AssetId            int
	Asset              Asset
	RelationshipTypeId int
	RelationshipType   AssetGroupRelationshipType
}

// Search criteria for locating asset group relationships.
type AssetGroupRelationshipSearchCriteria struct {
	rdb.Pagination
}
