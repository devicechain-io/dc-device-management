/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package schema

import (
	v1 "github.com/devicechain-io/dc-device-management/schema/v1"
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Drop all tables from the list.
func dropTables(tx *gorm.DB, tables []string) error {
	for table := range tables {
		err := tx.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}

// Creates the initial schema migration for this functional area.
func NewInitialSchema() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20220101000000",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&v1.Device{}, &v1.DeviceType{}, &v1.DeviceRelationshipType{}, &v1.DeviceRelationship{},
				&v1.DeviceGroup{}, &v1.DeviceGroupRelationshipType{}, &v1.DeviceGroupRelationship{},
				&v1.DeviceAssignmentStatus{}, &v1.DeviceAssignment{},

				&v1.AssetType{}, &v1.Asset{}, &v1.AssetRelationshipType{}, &v1.AssetRelationship{}, &v1.AssetGroup{},
				&v1.AssetGroupRelationshipType{}, &v1.AssetGroupRelationship{},

				&v1.CustomerType{}, &v1.Customer{}, &v1.CustomerRelationshipType{},
				&v1.CustomerRelationship{}, &v1.CustomerGroup{}, &v1.CustomerGroupRelationshipType{}, &v1.CustomerGroupRelationship{},

				&v1.AreaType{}, &v1.Area{}, &v1.AreaRelationshipType{}, &v1.AreaRelationship{}, &v1.AreaGroup{},
				&v1.AreaGroupRelationshipType{}, &v1.AreaGroupRelationship{})
		},
		Rollback: func(tx *gorm.DB) error {
			err := dropTables(tx, []string{"devices", "device_types", "device_relationship_types", "device_relationships",
				"device_groups", "device_group_relationship_types", "device_group_relationships", "device_assignment_statuses",
				"device_assignments"})
			if err != nil {
				return err
			}

			err = dropTables(tx, []string{"assets", "asset_types", "asset_relationship_types", "asset_relationships",
				"asset_groups", "asset_group_relationship_types", "asset_group_relationships"})
			if err != nil {
				return err
			}

			err = dropTables(tx, []string{"customers", "customer_types", "customer_relationship_types", "customer_relationships",
				"customer_groups", "customer_group_relationship_types", "customer_group_relationships"})
			if err != nil {
				return err
			}

			err = dropTables(tx, []string{"areas", "area_types", "area_relationship_types", "area_relationships",
				"area_groups", "area_group_relationship_types", "area_group_relationships"})
			if err != nil {
				return err
			}

			return nil
		},
	}
}
