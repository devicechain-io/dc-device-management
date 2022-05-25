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
				rdb.MetadataEntity

				Devices []Device
			}

			// Represents a device.
			type Device struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity

				DeviceTypeId uint
				DeviceType   *DeviceType

				Assignments []DeviceAssignment
			}

			// Metadata indicating a relationship between devices.
			type DeviceRelationshipType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity
			}

			// Captures a relationship between devices.
			type DeviceRelationship struct {
				gorm.Model
				rdb.MetadataEntity
				SourceDeviceId     uint
				SourceDevice       Device
				TargetDeviceId     uint
				TargetDevice       Device
				RelationshipTypeId uint
				RelationshipType   DeviceRelationshipType
			}

			// Represents a group of devices.
			type DeviceGroup struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.BrandedEntity
				rdb.MetadataEntity
			}

			// Metadata indicating a relationship between device and group.
			type DeviceGroupRelationshipType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity
			}

			// Represents a device-to-group relationship.
			type DeviceGroupRelationship struct {
				gorm.Model
				rdb.MetadataEntity
				DeviceGroupId      uint
				DeviceGroup        DeviceGroup
				DeviceId           uint
				Device             Device
				RelationshipTypeId uint
				RelationshipType   DeviceGroupRelationshipType
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

			// Represents an asset.
			type Asset struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity

				AssetTypeId uint
				AssetType   *AssetType
			}

			// Metadata indicating a relationship between assets.
			type AssetRelationshipType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity
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

			// Represents a group of assets.
			type AssetGroup struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.BrandedEntity
				rdb.MetadataEntity
			}

			// Metadata indicating a relationship between asset and group.
			type AssetGroupRelationshipType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity
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
			}

			// Represents a customer type.
			type CustomerType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.BrandedEntity
				rdb.MetadataEntity

				Customers []Customer
			}

			// Represents a customer.
			type Customer struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity

				CustomerTypeId uint
				CustomerType   *CustomerType
			}

			// Metadata indicating a relationship between customers.
			type CustomerRelationshipType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity
			}

			// Captures a relationship between customers.
			type CustomerRelationship struct {
				gorm.Model
				rdb.MetadataEntity
				SourceCustomerId   uint
				SourceCustomer     Customer
				TargetCustomerId   uint
				TargetCustomer     Customer
				RelationshipTypeId uint
				RelationshipType   CustomerRelationshipType
			}

			// Represents a group of customers.
			type CustomerGroup struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.BrandedEntity
				rdb.MetadataEntity
			}

			// Metadata indicating a relationship between customer and group.
			type CustomerGroupRelationshipType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity
			}

			// Represents a customer-to-group relationship.
			type CustomerGroupRelationship struct {
				gorm.Model
				rdb.MetadataEntity
				CustomerGroupId    uint
				CustomerGroup      CustomerGroup
				CustomerId         uint
				Customer           Customer
				RelationshipTypeId uint
				RelationshipType   CustomerGroupRelationshipType
			}

			// Represents an area type.
			type AreaType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.BrandedEntity
				rdb.MetadataEntity

				Areas []Area
			}

			// Represents an area.
			type Area struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity

				AreaTypeId uint
				AreaType   *AreaType
			}

			// Metadata indicating a relationship between areas.
			type AreaRelationshipType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity
			}

			// Captures a relationship between areas.
			type AreaRelationship struct {
				gorm.Model
				rdb.MetadataEntity
				SourceAreaId       uint
				SourceArea         Area
				TargetAreaId       uint
				TargetArea         Area
				RelationshipTypeId uint
				RelationshipType   AreaRelationshipType
			}

			// Represents a group of areas.
			type AreaGroup struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.BrandedEntity
				rdb.MetadataEntity
			}

			// Metadata indicating a relationship between area and group.
			type AreaGroupRelationshipType struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity
			}

			// Represents a area-to-group relationship.
			type AreaGroupRelationship struct {
				gorm.Model
				rdb.MetadataEntity
				AreaGroupId        uint
				AreaGroup          AreaGroup
				AreaId             uint
				Area               Area
				RelationshipTypeId uint
				RelationshipType   AreaGroupRelationshipType
			}

			// Metadata indicating a relationship between devices.
			type DeviceAssignmentStatus struct {
				gorm.Model
				rdb.TokenReference
				rdb.NamedEntity
				rdb.MetadataEntity
			}

			// Provides context for device.
			type DeviceAssignment struct {
				gorm.Model
				rdb.TokenReference
				rdb.MetadataEntity
				DeviceId                 uint
				Device                   Device
				DeviceGroupId            *uint
				DeviceGroup              DeviceGroup
				AssetId                  *uint
				Asset                    Asset
				AssetGroupId             *uint
				AssetGroup               AssetGroup
				CustomerId               *uint
				Customer                 Customer
				CustomerGroupId          *uint
				CustomerGroup            CustomerGroup
				AreaId                   *uint
				Area                     Area
				AreaGroupId              *uint
				AreaGroup                AreaGroup
				DeviceAssignmentStatusId *uint
				DeviceAssignmentStatus   DeviceAssignmentStatus
				Active                   bool
			}

			return tx.AutoMigrate(&Device{}, &DeviceType{}, &DeviceRelationshipType{}, &DeviceRelationship{},
				&DeviceGroup{}, &DeviceGroupRelationshipType{}, &DeviceGroupRelationship{},
				&DeviceAssignmentStatus{}, &DeviceAssignment{},

				&AssetType{}, &Asset{}, &AssetRelationshipType{}, &AssetRelationship{}, &AssetGroup{},
				&AssetGroupRelationshipType{}, &AssetGroupRelationship{},

				&CustomerType{}, &Customer{}, &CustomerRelationshipType{},
				&CustomerRelationship{}, &CustomerGroup{}, &CustomerGroupRelationshipType{}, &CustomerGroupRelationship{},

				&AreaType{}, &Area{}, &AreaRelationshipType{}, &AreaRelationship{}, &AreaGroup{},
				&AreaGroupRelationshipType{}, &AreaGroupRelationship{})
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
			err = tx.Migrator().DropTable("asset_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("assets")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("asset_relationship_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("asset_relationships")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("asset_groups")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("asset_group_relationship_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("asset_group_relationships")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("customer_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("customer")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("customer_relationship_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("customer_relationships")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("customer_groups")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("customer_group_relationship_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("customer_group_relationships")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("area_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("area")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("area_relationship_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("area_relationships")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("area_groups")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("area_group_relationship_types")
			if err != nil {
				return err
			}
			err = tx.Migrator().DropTable("area_group_relationships")
			if err != nil {
				return err
			}
			return nil
		},
	}
}
