/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var (
	Migrations = []*gormigrate.Migration{
		{
			ID: "20220322151200",
			Migrate: func(tx *gorm.DB) error {
				type Device struct {
					gorm.Model
					Name string
				}

				return tx.AutoMigrate(&Device{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("devices")
			},
		},
	}
)
