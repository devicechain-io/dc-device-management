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

// Results for asset type search.
type AssetTypeSearchResults struct {
	Results    []AssetType
	Pagination rdb.SearchResultsPagination
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

	AssetTypeId uint
	AssetType   *AssetType
}

// Search criteria for locating assets.
type AssetSearchCriteria struct {
	rdb.Pagination
	AssetTypeToken *string
}

// Results for asset search.
type AssetSearchResults struct {
	Results    []Asset
	Pagination rdb.SearchResultsPagination
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

// Results for asset relationship type search.
type AssetRelationshipTypeSearchResults struct {
	Results    []AssetRelationshipType
	Pagination rdb.SearchResultsPagination
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
	SourceAssetId      uint
	SourceAsset        Asset
	TargetAssetId      uint
	TargetAsset        Asset
	RelationshipTypeId uint
	RelationshipType   AssetRelationshipType
}

// Search criteria for locating asset relationships.
type AssetRelationshipSearchCriteria struct {
	rdb.Pagination
}

// Results for asset relationship search.
type AssetRelationshipSearchResults struct {
	Results    []AssetRelationship
	Pagination rdb.SearchResultsPagination
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

// Results for asset group search.
type AssetGroupSearchResults struct {
	Results    []AssetGroup
	Pagination rdb.SearchResultsPagination
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

// Search criteria for locating asset group relationship types.
type AssetGroupRelationshipTypeSearchCriteria struct {
	rdb.Pagination
}

// Results for asset group relationship type search.
type AssetGroupRelationshipTypeSearchResults struct {
	Results    []AssetGroupRelationshipType
	Pagination rdb.SearchResultsPagination
}

// Data required to create an asset group relationship.
type AssetGroupRelationshipCreateRequest struct {
	AssetGroup       string
	Asset            string
	RelationshipType string
	Metadata         *string
}

// Represents a asset-to-group relationship.
type AssetGroupRelationship struct {
	gorm.Model
	rdb.MetadataEntity
	AssetGroupId       uint
	AssetGroup         AssetGroup
	AssetId            uint
	Asset              Asset
	RelationshipTypeId uint
	RelationshipType   AssetGroupRelationshipType
}

// Search criteria for locating asset group relationships.
type AssetGroupRelationshipSearchCriteria struct {
	rdb.Pagination
}

// Results for asset group relationship search.
type AssetGroupRelationshipSearchResults struct {
	Results    []AssetGroupRelationship
	Pagination rdb.SearchResultsPagination
}
