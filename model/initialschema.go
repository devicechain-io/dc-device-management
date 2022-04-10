/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"github.com/devicechain-io/dc-microservice/rdb"
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Creates the initial schema migration for this functional area.
func NewInitialSchema() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20220101000000",
		Migrate: func(tx *gorm.DB) error {
			// Represents a device type.
			type DeviceType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.BrandedEntity
				Devices []Device
			}

			// Represents a device.
			type Device struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				DeviceTypeId int
				DeviceType   DeviceType
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

			// Metadata indicating a relationship between device and group.
			type DeviceGroupRelationshipType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
			}

			// Represents a device-to-group relationship.
			type DeviceGroupRelationship struct {
				gorm.Model
				DeviceGroupId      int
				DeviceGroup        DeviceGroup
				DeviceId           int
				Device             Device
				RelationshipTypeId int
				RelationshipType   DeviceGroupRelationshipType
			}

			// Represents an asset type.
			type AssetType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.BrandedEntity

				Assets []Asset
			}

			// Represents an asset.
			type Asset struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity

				AssetTypeId int
				AssetType   *AssetType
			}

			// Metadata indicating a relationship between assets.
			type AssetRelationshipType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
			}

			// Captures a relationship between assets.
			type AssetRelationship struct {
				gorm.Model
				SourceAssetId      int
				SourceAsset        Asset
				TargetAssetId      int
				TargetAsset        Asset
				RelationshipTypeId int
				RelationshipType   AssetRelationshipType
			}

			// Represents a group of assets.
			type AssetGroup struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.BrandedEntity
			}

			// Metadata indicating a relationship between asset and group.
			type AssetGroupRelationshipType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
			}

			// Represents a asset-to-group relationship.
			type AssetGroupRelationship struct {
				gorm.Model
				AssetGroupId       int
				AssetGroup         AssetGroup
				AssetId            int
				Asset              Asset
				RelationshipTypeId int
				RelationshipType   AssetGroupRelationshipType
			}

			return tx.AutoMigrate(&Device{}, &DeviceType{}, &DeviceRelationshipType{}, &DeviceRelationship{},
				&DeviceGroup{}, &DeviceGroupRelationshipType{}, &DeviceGroupRelationship{}, &AssetType{},
				&Asset{}, &AssetRelationshipType{}, &AssetRelationship{}, &AssetGroup{}, &AssetGroupRelationshipType{},
				&AssetGroupRelationship{})
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Migrator().DropTable("device_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("devices")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("device_relationship_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("device_relationships")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("device_groups")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("device_group_relationship_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("device_group_relationships")
			if err != nil {
				return err
			}
			return nil
		},
	}
}
