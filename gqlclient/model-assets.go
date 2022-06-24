/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package gqlclient

// Asset type entity.
type IAssetType interface {
	IModel
	ITokenReference
	INamedEntity
	IBrandedEntity
	IMetadataEntity
}

// Asset entity.
type IAsset interface {
	IModel
	ITokenReference
	INamedEntity
	IMetadataEntity
	GetAssetType() DefaultAssetAssetType
}

// Asset relationship type entity.
type IAssetRelationshipType interface {
	IModel
	ITokenReference
	INamedEntity
	IMetadataEntity
}

// Asset relationship entity.
type IAssetRelationship interface {
	IModel
	ITokenReference
	IMetadataEntity
	GetSourceAsset() DefaultAssetRelationshipSourceAsset
	GetTargetAsset() DefaultAssetRelationshipTargetAsset
	GetRelationshipType() DefaultAssetRelationshipRelationshipTypeAssetRelationshipType
}

// Asset group entity.
type IAssetGroup interface {
	IModel
	ITokenReference
	INamedEntity
	IBrandedEntity
	IMetadataEntity
}

// Asset group relationship type entity.
type IAssetGroupRelationshipType interface {
	IModel
	ITokenReference
	INamedEntity
	IMetadataEntity
}

// Asset group relationship entity.
type IAssetGroupRelationship interface {
	IModel
	ITokenReference
	IMetadataEntity
	GetAssetGroup() DefaultAssetGroupRelationshipAssetGroup
	GetAsset() DefaultAssetGroupRelationshipAsset
	GetRelationshipType() DefaultAssetGroupRelationshipRelationshipTypeAssetGroupRelationshipType
}
